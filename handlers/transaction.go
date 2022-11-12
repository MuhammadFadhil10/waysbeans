package handlers

import (
	"encoding/json"
	"net/http"
	transactiondto "waysbeans/dto/transaction"
	"waysbeans/helper"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := transactiondto.CreateTransactionRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	transaction := models.Transaction{
		UserID:   userId,
		Status:   request.Status,
		Products: request.Products,
	}

	var transactionResponse models.Transaction
	var err error
	transactionResponse, err = h.TransactionRepository.CreateTransaction(transaction)

	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	helper.ResponseHelper(w, nil, transactionResponse, 0)

}
