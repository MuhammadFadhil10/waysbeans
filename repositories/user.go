package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetProfile(user models.User, userID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProfile(user models.User, userID int) (models.User, error) {
	err := r.db.First(&user, userID).Error

	return user, err
}
