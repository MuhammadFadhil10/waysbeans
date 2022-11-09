package handlers

import (
	"encoding/json"
	"net/http"
	transactiondto "waysbeans/dto/transaction"
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

	transactionRequest := transactiondto.CreateTransactionRequest{}
	json.NewDecoder(r.Body).Decode(&transactionRequest)

}
