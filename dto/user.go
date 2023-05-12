package dto

import (
	"hacktiv8-msib-final-project-3/entity"
	"time"
)

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (r *RegisterRequest) ToEntity() *entity.User {
	return &entity.User{
		FullName: r.FullName,
		Email:    r.Email,
		Password: r.Password,
		Role:     "member",
	}
}

type RegisterResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
