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

type GetAllCategoriesResponse struct {
	ID        uint       `json:"id"`
	Type      string     `json:"type"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Tasks     []TaskData `json:"tasks"`
}

type UpdateCategoryRequest CreateCategoryRequest

func (c *UpdateCategoryRequest) ToEntity() *entity.Category {
	return &entity.Category{
		Type: c.Type,
	}
}

type UpdateCategoryResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCategoryResponse struct {
	Message string `json:"message"`
}
