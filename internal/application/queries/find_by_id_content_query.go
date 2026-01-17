package queries

import (
	"context"

	"github.com/google/uuid"
	"github.com/thalesdev/dang/internal/domain/entities"
	"github.com/thalesdev/dang/internal/domain/repositories"
)

type FindByIdContentQuery struct {
	id uuid.UUID
}

func NewFindByIdContentQuery(id uuid.UUID) FindByIdContentQuery {
	return FindByIdContentQuery{
		id: id,
	}
}

type FindByIdContentHandler struct {
	repository repositories.ContentRepository
}

func NewFindByIdContentHandler(repository repositories.ContentRepository) *FindByIdContentHandler {
	return &FindByIdContentHandler{
		repository: repository,
	}
}

func (f *FindByIdContentHandler) Execute(ctx context.Context, query FindByIdContentQuery) (*entities.Content, error) {
	content, err := f.repository.FindById(ctx, query.id)
	if err != nil {
		return nil, err
	}
	return content, nil
}
