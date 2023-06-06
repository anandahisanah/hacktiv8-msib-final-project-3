package dto

import (
	"time"

	"hacktiv8-msib-final-project-3/entity"
)

type TaskData struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_Id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  uint   `json:"category_id" `
}

func (t *CreateTaskRequest) ToEntity() *entity.Task {
	return &entity.Task{
		Title:       t.Title,
		Description: t.Description,
		CategoryID:  t.CategoryID,
	}
}

type CreateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetAllTasksResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User        UserData  `json:"user"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (t *UpdateTaskRequest) ToEntity() *entity.Task {
	return &entity.Task{
		Title:       t.Title,
		Description: t.Description,
	}
}

type UpdateTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTaskStatusRequest struct {
	Status bool `json:"status"`
}

type UpdateTaskCategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type DeleteTaskResponse struct {
	Message string `json:"message"`
}
