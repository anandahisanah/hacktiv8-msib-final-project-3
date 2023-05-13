package service

import (
	"hacktiv8-msib-final-project-3/dto"
	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/repository/userrepository"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(userRepo userrepository.UserRepository) UserService {
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

func (u *userService) UpdateUser(user *entity.User, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {
	newUser := payload.ToEntity()

	updatedUser, err := u.userRepo.UpdateUser(user, newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		ID:        updatedUser.ID,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}
