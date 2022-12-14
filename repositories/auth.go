package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GetByEmail(user models.User, email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	LoginUser(user models.User, email string) (models.User, error)
	GetAuthProfile(user models.User, userID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByEmail(user models.User, email string) (models.User, error) {
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) LoginUser(user models.User, email string) (models.User, error) {
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) GetAuthProfile(user models.User, userID int) (models.User, error) {
	err := r.db.First(&user, userID).Error

	return user, err
}
