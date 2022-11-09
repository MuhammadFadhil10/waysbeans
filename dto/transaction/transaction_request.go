package transactiondto

import "waysbeans/models"

type CreateTransactionRequest struct {
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address"`
	Attachment string  `json:"attachment"`
	Products   []models.ProductTransactionResponse `json:"products"`
}