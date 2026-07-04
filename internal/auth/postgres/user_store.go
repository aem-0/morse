// internal/auth/postgres/user_store.go
package postgres

import (
	"context"
	"fmt"

	"morse/internal/auth/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore struct {
	pool *pgxpool.Pool
}

var _ domain.UserStore = (*UserStore)(nil)

func NewUserStore(pool *pgxpool.Pool) *UserStore {
	return &UserStore{pool: pool}
}

func (s *UserStore) Save(ctx context.Context, user domain.User) error {
	query := `
        INSERT INTO users (id, username, is_active, created_at, updated_at)
        VALUES (@id, @username, @isActive, @createdAt, @updatedAt)
    `
	args := pgx.NamedArgs{
		"id":        user.ID,
		"username":  user.UserName,
		"isActive":  user.IsActive,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	}

	_, err := s.pool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("UserStore.Save: %w", err)
	}
	return nil
}

func (s *UserStore) Get(ctx context.Context, userName string) (*domain.User, error) {
	query := `
        SELECT id, username, is_active, last_authenticated_at, created_at, updated_at
        FROM users
        WHERE username = @username AND is_active = true
    `
	args := pgx.NamedArgs{
		"username": userName,
	}

	rows, err := s.pool.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("UserStore.Get: %w", err)
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("UserStore.Get: %w", err)
	}

	return &user, nil
}

func (s *UserStore) UpdateLastAuthenticated(ctx context.Context, userID uuid.UUID) error {
    query := `
        UPDATE users 
        SET last_authenticated_at = NOW(),
            updated_at = NOW()
        WHERE id = @id
    `
    args := pgx.NamedArgs{
        "id": userID,
    }

    _, err := s.pool.Exec(ctx, query, args)
    if err != nil {
        return fmt.Errorf("UserStore.UpdateLastAuthenticated: %w", err)
    }
    return nil
}