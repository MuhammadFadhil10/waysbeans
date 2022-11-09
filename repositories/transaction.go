package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(t models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(t models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&t).Error

	return t, err
}
