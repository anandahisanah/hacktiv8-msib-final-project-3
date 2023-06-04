package dto

import (
	"time"

	"hacktiv8-msib-final-project-3/entity"
)

// ==============================================================

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

// ==========================================

type LoginRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"jwt"`
}

// =============================================

type UpdateUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
}

func (r *UpdateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		FullName: r.FullName,
		Email:    r.Email,
	}
}

type UpdateUserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ===============================

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type UserData struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
