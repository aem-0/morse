package domain

import (
	"time"

	"github.com/google/uuid"
)

type PublicKey struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Material       string
	KeyFingerprint string
	IsActive       bool
	CreatedAt      time.Time
	RevokedAt      time.Time
}
