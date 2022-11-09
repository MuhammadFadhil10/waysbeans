package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart models.Cart) (models.Cart, error)
	GetCart(cart models.Cart, ID int) (models.Cart, error)
	GetCarts(carts []models.Cart) ([]models.Cart, error)
	UpdateCartQty(cart models.Cart, ID int) (models.Cart, error)
	DeleteCartByID(cart models.Cart, ID int) (models.Cart, error)
	DeleteCartByUser(cart models.Cart, userID int) error
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

func (r *repository) GetCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Preload("User").Preload("Products").First(&cart).Error

	return cart, err
}

func (r *repository) UpdateCartQty(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Model(&cart).Where("id=?", ID).Preload("User").Preload("Products").Updates(&cart).Error

	return cart, err
}

func (r *repository) DeleteCartByID(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart, "id=?", ID).Preload("Products").Preload("User").Error
	return cart, err
}

func (r *repository) DeleteCartByUser(cart models.Cart, userID int) error {
	err := r.db.Preload("Products").Preload("User").Delete(&cart, "user_id=?", userID).Error
	return err
}
