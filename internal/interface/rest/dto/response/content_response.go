package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/thalesdev/dang/internal/domain/entities"
)

type ContentResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ContentListResponse struct {
	Items []ContentResponse `json:"items"`
}

func ToContentResponse(content entities.Content) ContentResponse {
	return ContentResponse{
		ID:          content.ID,
		Name:        content.Name,
		Description: content.Description,
		CreatedAt:   content.CreatedAt,
		UpdatedAt:   content.UpdatedAt,
	}
}

func ToContentResponseList(content []entities.Content) []ContentResponse {
	res := make([]ContentResponse, len(content))
	for i := range content {
		res[i] = ToContentResponse(content[i])
	}
	return res
}
