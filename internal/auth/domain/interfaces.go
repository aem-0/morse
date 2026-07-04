package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserStore interface {
	Save(ctx context.Context, user User) error
	Get(ctx context.Context, userName string) (*User, error)
	    UpdateLastAuthenticated(ctx context.Context, userID uuid.UUID) error 
}

type PublicKeyStore interface {
	Save(ctx context.Context, key PublicKey) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*PublicKey, error)
	GetByFingerprint(ctx context.Context, fingerprint string) (*PublicKey, error)
	UpdateLastUsed(ctx context.Context, keyID uuid.UUID) error
}
