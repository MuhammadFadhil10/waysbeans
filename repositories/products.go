package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	GetProducts(products []models.Products) ([]models.Products, error)
}

func RepositoryProducts(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts(products []models.Products) ([]models.Products, error) {
	err := r.db.Find(&products).Error

	return products, err
}
