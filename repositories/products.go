package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	GetProducts(products []models.Products) ([]models.Products, error)
	GetProductById(product models.Products, productId int) (models.Products, error)
	CreateProduct(product models.Products) (models.Products, error)
}

func RepositoryProducts(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts(products []models.Products) ([]models.Products, error) {
	err := r.db.Find(&products).Error

	return products, err
}
func (r *repository) GetProductById(product models.Products, productId int) (models.Products, error) {
	err := r.db.First(&product, productId).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Products) (models.Products, error) {
	err := r.db.Create(&product).Error
	return product, err
}
