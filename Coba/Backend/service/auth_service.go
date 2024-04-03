package service

import (
	"backend/entity"
	"backend/errorHandler"
	"backend/repository"
	"log"
)

type AuthService interface {
	Register(req *entity.User) (*entity.RegisterResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *entity.User) (*entity.RegisterResponse, error) {
	var data entity.RegisterResponse
	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {
		log.Print("username already exists")
		return nil, errorHandler.ConflictError("Username Already Exists")
	}
	user := entity.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	data = entity.RegisterResponse{
		Username: user.Username,
		Name:     user.Name,
	}

	if err := s.repository.Register(&user); err != nil {
		return nil, errorHandler.InternalServerError(err.Error())
	}
	return &data, nil
}
