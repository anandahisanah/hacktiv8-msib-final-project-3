package dto

import (
	"hacktiv8-msib-final-project-3/entity"
	"time"
)

type CreateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

func (c *CreateCategoryRequest) ToEntity() *entity.Category {
	return &entity.Category{
		Type: c.Type,
	}
}

type CreateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
