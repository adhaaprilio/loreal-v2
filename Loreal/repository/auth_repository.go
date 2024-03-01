package repository

import (
	"loreal/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	UsernameExist(username string) bool
	Register(req *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entity.User) error {
	err := r.db.Create(&user).Error
	return err
}
func (r *authRepository) UsernameExist(username string) bool {
	var user entity.User
	err := r.db.First(&user, "username=?", username).Error

	return err == nil
}

func (r *authRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "username = ?", username).Error

	return &user, err
}
