package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/thalesdev/dang/internal/domain/entities"
)

type ContentFilter struct {
	Title *string
}

type ContentRepository interface {
	FindAll(ctx context.Context, filters ContentFilter) ([]entities.Content, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.Content, error)
}
