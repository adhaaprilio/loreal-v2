package service

import (
	"loreal/dto"
	"loreal/entity"
	"loreal/errorHandler"
	"loreal/helper"
	"loreal/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthSerivce(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {
		return &errorHandler.BadRequestError{Message: "Username already exist"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorHandler.BadRequestError{Message: "Please enter the same password"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorHandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:     req.Name,
		Username: req.Username,
		Password: passwordHash,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorHandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByUsername(req.Username)

	if err != nil {
		return nil, &errorHandler.NotFoundError{Message: "Wrong username or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorHandler.NotFoundError{Message: "Wrong username or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
