package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart models.Cart) (models.Cart, error)
	GetCarts(carts []models.Cart) ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddToCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Preload("User").Preload("Products").Error

	return cart, err
}

func (r *repository) GetCarts(carts []models.Cart) ([]models.Cart, error) {
	err := r.db.Preload("User").Preload("Products").Find(&carts).Error

	return carts, err
}
