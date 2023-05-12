package service

import (
	"hacktiv8-msib-final-project-3/dto"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/repository/user_repository"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo}
}

// ===========================================================================

func (u *userService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	user := payload.ToEntity()
	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	registeredUser, err := u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:        registeredUser.ID,
		FullName:  registeredUser.FullName,
		Email:     registeredUser.Email,
		CreatedAt: registeredUser.CreatedAt,
	}

	return response, nil
}

func (u *userService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	response := &dto.LoginResponse{Token: token}

	return response, nil
}
