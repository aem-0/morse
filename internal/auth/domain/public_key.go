package domain

import (
	"time"

	"github.com/google/uuid"
)

type PublicKey struct {
	ID             uuid.UUID  `db:"id"`
	UserID         uuid.UUID  `db:"user_id"`
	PublicKey      []byte     `db:"public_key"`
	KeyFingerprint string     `db:"key_fingerprint"`
	KeyType        string     `db:"key_type"`
	DeviceName     string     `db:"device_name"`
	IsActive       bool       `db:"is_active"`
	CreatedAt      time.Time  `db:"created_at"`
	LastUsedAt     *time.Time `db:"last_used_at"`
	RevokedAt      *time.Time `db:"revoked_at"`
}
