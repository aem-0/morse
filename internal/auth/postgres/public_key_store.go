// internal/auth/postgres/public_key_store.go
package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"morse/internal/auth/domain"
)

type PublicKeyStore struct {
	pool *pgxpool.Pool
}

// compile time check
var _ domain.PublicKeyStore = (*PublicKeyStore)(nil)

func NewPublicKeyStore(pool *pgxpool.Pool) *PublicKeyStore {
	return &PublicKeyStore{pool: pool}
}

func (s *PublicKeyStore) Save(ctx context.Context, key domain.PublicKey) error {
	query := `
		INSERT INTO public_keys (
			id,
			user_id,
			public_key,
			key_fingerprint,
			key_type,
			device_name,
			is_active,
			created_at
		) VALUES (
			@id,
			@userID,
			@publicKey,
			@keyFingerprint,
			@keyType,
			@deviceName,
			@isActive,
			@createdAt
		)
	`

	args := pgx.NamedArgs{
		"id":             key.ID,
		"userID":         key.UserID,
		"publicKey":      key.PublicKey,
		"keyFingerprint": key.KeyFingerprint,
		"keyType":        key.KeyType,
		"deviceName":     key.DeviceName,
		"isActive":       key.IsActive,
		"createdAt":      key.CreatedAt,
	}

	_, err := s.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("PublicKeyStore.Save: %w", err)
	}
	return nil
}

func (s *PublicKeyStore) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.PublicKey, error) {
	query := `
		SELECT
			id,
			user_id,
			public_key,
			key_fingerprint,
			key_type,
			device_name,
			is_active,
			created_at,
			last_used_at,
			revoked_at
		FROM public_keys
		WHERE user_id = @userID
		  AND is_active = true
	`

	args := pgx.NamedArgs{
		"userID": userID,
	}

	rows, err := s.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("PublicKeyStore.GetByUserID: %w", err)
	}

	keys, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.PublicKey])
	if err != nil {
		return nil, fmt.Errorf("PublicKeyStore.GetByUserID scan: %w", err)
	}

	// convert to slice of pointers
	result := make([]*domain.PublicKey, len(keys))
	for i := range keys {
		result[i] = &keys[i]
	}

	return result, nil
}

func (s *PublicKeyStore) GetByFingerprint(ctx context.Context, fingerprint string) (*domain.PublicKey, error) {
	query := `
		SELECT
			id,
			user_id,
			public_key,
			key_fingerprint,
			key_type,
			device_name,
			is_active,
			created_at,
			last_used_at,
			revoked_at
		FROM public_keys
		WHERE key_fingerprint = @fingerprint
		  AND is_active = true
	`

	args := pgx.NamedArgs{
		"fingerprint": fingerprint,
	}

	rows, err := s.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("PublicKeyStore.GetByFingerprint: %w", err)
	}

	key, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.PublicKey])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // key not found, not an error
		}
		return nil, fmt.Errorf("PublicKeyStore.GetByFingerprint scan: %w", err)
	}

	return &key, nil
}

func (s *PublicKeyStore) UpdateLastUsed(ctx context.Context, keyID uuid.UUID) error {
	query := `
        UPDATE public_keys
        SET last_used_at = NOW()
        WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": keyID,
	}

	_, err := s.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("PublicKeyStore.UpdateLastUsed: %w", err)
	}
	return nil
}
