package handlers

import (
	"encoding/json"
	"net/http"
	transactiondto "waysbeans/dto/transaction"
	"waysbeans/models"
	"waysbeans/repositories"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := transactiondto.CreateTransactionRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	transaction := models.Transaction{
		UserID: request.UserID,
		Status: request.Status,
		Products: request.Products,
	}
	h.TransactionRepository.CreateTransaction(transaction)

}
