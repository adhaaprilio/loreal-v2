package repository

import (
	"ProjectSprint/models"
	"database/sql"
	"fmt"
)

type AuthRepository interface {
	UsernameExist(username string) bool
	Register(req *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *models.User) error {
	query := `INSERT INTO users(username, name, password)
				VALUES ($1,$2,$3);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Name, user.Password)
	return err
}

func (r *authRepository) UsernameExist(username string) bool {
	var exists bool
	query := `SELECT EXISTS(
		SELECT 1
		FROM users
		WHERE username = $1
	)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&exists)
	if err != nil {
		return false
	}

	return exists

}

func (r *authRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT username, name, password
		FROM users
		WHERE username = $1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.Name, &user.Password)
	fmt.Print(err)
	if err != nil {
		return nil, err
	}
	return &user, err
}
