package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID  `db:"id"`
	UserName            string     `db:"username"`
	IsActive            bool       `db:"is_active"`
	LastAuthenticatedAt *time.Time `db:"last_authenticated_at"`
	CreatedAt           time.Time  `db:"created_at"`
	UpdatedAt           time.Time  `db:"updated_at"`
}
