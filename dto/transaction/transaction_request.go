package transactiondto

import "waysbeans/models"

type CreateTransactionRequest struct {
	UserID   int               `Json:"userId"`
	Status   string            `json:"status"`
	Products []models.Products `json:"products"`
}
