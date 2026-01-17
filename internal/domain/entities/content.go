package entities

import (
	"time"

	"github.com/google/uuid"
)

type Content struct {
	ID          uuid.UUID
	Name        string
	Description string

	CreatedAt *time.Time
	UpdatedAt *time.Time
}
