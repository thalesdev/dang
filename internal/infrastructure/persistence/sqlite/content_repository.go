package sqlite

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/thalesdev/dang/internal/domain/entities"
	"github.com/thalesdev/dang/internal/domain/repositories"
	"github.com/thalesdev/dang/internal/infrastructure/persistence/sqlite/models"
)

type ContentRepository struct {
	db *sqlx.DB
}

var _ repositories.ContentRepository = (*ContentRepository)(nil)

func NewContentRepository(db *sqlx.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

func (c *ContentRepository) FindAll(ctx context.Context, filter repositories.ContentFilter) ([]entities.Content, error) {
	var m []models.ContentModel
	query := "select id, name, description, created_at, updated_at from content where 1 = 1"

	var args []interface{}

	if filter.Title != nil {
		query += " and name like ?"
		args = append(args, fmt.Sprintf("%%%s%%", *filter.Title))
	}

	err := c.db.SelectContext(ctx, &m, query, args...)
	if err != nil {
		return nil, err
	}

	entityList := make([]entities.Content, len(m))
	for idx, item := range m {
		entityList[idx] = *item.ToDomain()
	}

	return entityList, nil
}

func (c *ContentRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.Content, error) {
	var m models.ContentModel
	query := "select id, name, description, created_at, updated_at from content where id = ?"

	err := c.db.GetContext(ctx, &m, query, id.String())
	if err != nil {
		return nil, err
	}

	return m.ToDomain(), nil
}
