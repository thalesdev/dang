package models

import (
	"github.com/google/uuid"
	"github.com/thalesdev/dang/internal/domain/entities"
)

type ContentModel struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`

	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func FromDomain(content entities.Content) *ContentModel {

	return &ContentModel{
		ID:          content.ID,
		Name:        content.Name,
		Description: content.Description,
		CreatedAt:   content.CreatedAt.String(),
		UpdatedAt:   content.UpdatedAt.String(),
	}
}

func (cm *ContentModel) ToDomain() *entities.Content {

	return &entities.Content{
		ID:          cm.ID,
		Name:        cm.Name,
		Description: cm.Description,
		CreatedAt:   ParseTime(cm.CreatedAt),
		UpdatedAt:   ParseTime(cm.UpdatedAt),
	}
}
