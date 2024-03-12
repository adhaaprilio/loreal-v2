package services

import (
	"ProjectSprint/internal/dto"
	"ProjectSprint/internal/errorHandler"
	"ProjectSprint/internal/helper"
	"ProjectSprint/internal/repository"
	"ProjectSprint/models"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	var data dto.RegisterResponse
	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {
		return nil, &errorHandler.Error409{Message: "Username Already Exist"}
	}

	if isValidLength(req.Username, 5, 15) || isValidLength(req.Name, 5, 50) || isValidLength(req.Password, 5, 15) {
		return nil, &errorHandler.Error400{Message: "Password or username is too short or too long"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, &errorHandler.Error500{Message: err.Error()}
	}

	user := models.User{
		Username: req.Username,
		Name:     req.Name,
		Password: passwordHash,
	}

	if err := s.repository.Register(&user); err != nil {
		return nil, &errorHandler.Error500{Message: err.Error()}
	}
	token, err := helper.GenerateToken(&user)
	if err != nil {
		return nil, &errorHandler.Error500{Message: err.Error()}
	}

	data = dto.RegisterResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: token,
	}

	return &data, nil
}
func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse
	user, err := s.repository.GetUserByUsername(req.Username)

	if err != nil {
		return nil, &errorHandler.Error404{Message: "Username wrong"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorHandler.Error400{Message: "Username or password wrong"}
	}
	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorHandler.Error500{Message: err.Error()}
	}

	data = dto.LoginResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: token,
	}

	return &data, nil
}
func isValidLength(str string, minLength, maxLength int) bool {
	length := len(str)
	return length <= minLength && length >= maxLength
}
