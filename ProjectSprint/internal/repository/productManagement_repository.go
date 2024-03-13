package repository

import (
	"ProjectSprint/models"
	"database/sql"
)

type ProductManagementRepository interface {
	CreateProduct(req *models.Product) error
}

type productManagementRepository struct {
	db *sql.DB
}

func NewProductManagementRepository(db *sql.DB) *productManagementRepository {
	return &productManagementRepository{
		db: db,
	}
}

// func (r *productManagementRepository) CreateProduct(product *models.Product) error {
// 	// query := `INSERT INTO products`
// }
