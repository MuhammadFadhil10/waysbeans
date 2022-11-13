package repositories

import (
	"fmt"
	"waysbeans/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(t models.Transaction) (models.Transaction, error)
	GetAllTransactions() ([]models.Transaction, error)
	GetTransaction(transactionId int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	GetTransactionByUser(userID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(t models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&t).Error

	return t, err
}

func (r *repository) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := r.db.Preload("Products").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(transactionId int) (models.Transaction, error) {
	var transactionData models.Transaction
	err := r.db.First(&transactionData, transactionId).Error

	return transactionData, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Products").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Products").Preload("User").Order("created_at DESC").Find(&transactions, "user_id=?", userID).Error

	return transactions, err
}

func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	var carts []models.Cart

	r.db.Preload("Products").First(&transaction, ID)
	r.db.Preload("Products").Find(&carts, "user_id=?", transaction.UserID)

	fmt.Println("STATUS:", status)
	fmt.Println("STATUS TRANSACTION:", transaction.Status)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		for index, tp := range transaction.Products {
			// fmt.Println(tp.Qty)
			var product models.Products
			r.db.First(&product, tp.ID)
			product.Stock = product.Stock - carts[index].Qty
			r.db.Model(&product).Where("id=?", tp.ID).Updates(&product)
		}
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}
