package domain

import "github.com/google/uuid"

type PublicKeyRepository interface {
	Save(key PublicKey) error
	Get(userID uuid.UUID) ([]*PublicKey, error)
}
