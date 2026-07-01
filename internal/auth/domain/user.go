package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID
	UserName            string
	LastAuthenticatedAt time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	IsActive            bool
}
