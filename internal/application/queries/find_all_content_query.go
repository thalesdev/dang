package queries

import (
	"context"

	"github.com/thalesdev/dang/internal/domain/entities"
	"github.com/thalesdev/dang/internal/domain/repositories"
)

type FindAllContentQuery struct {
	title *string
}

func NewFindAllContentQuery(title *string) FindAllContentQuery {
	return FindAllContentQuery{title: title}
}

type FindAllContentHandler struct {
	repository repositories.ContentRepository
}

func (f *FindAllContentHandler) Execute(ctx context.Context, query FindAllContentQuery) ([]entities.Content, error) {
	items, err := f.repository.FindAll(ctx, repositories.ContentFilter{Title: query.title})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func NewFindAllContentHandler(repository repositories.ContentRepository) *FindAllContentHandler {
	return &FindAllContentHandler{repository: repository}
}
