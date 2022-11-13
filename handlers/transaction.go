package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	transactiondto "waysbeans/dto/transaction"
	"waysbeans/helper"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
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
		UserID:     userId,
		Status:     request.Status,
		Products:   request.Products,
		TotalPrice: request.TotalPrice,
	}

	// var transactionResponse models.Transaction
	// var err error
	newTransaction, err := h.TransactionRepository.CreateTransaction(transaction)

	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	// Create Unique Transaction Id
	// var TransIdIsMatch = false
	// var TransactionId int
	// var transactionData models.Transaction
	// for !TransIdIsMatch {
	// 	TransactionId = userId + newTransaction.ID + rand.Intn(10000) - rand.Intn(100)
	// 	transactionData, _ = h.TransactionRepository.GetTransaction(TransactionId)
	// 	if transactionData.ID == 0 {
	// 		TransIdIsMatch = true
	// 	}
	// }

	// Request token transaction from midtrans
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(newTransaction.ID),
			GrossAmt: int64(newTransaction.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: newTransaction.User.FullName,
			Email: newTransaction.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	helper.ResponseHelper(w, nil, snapResp, 0)

}

func (h *handlerTransaction) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.GetAllTransactions()
	fmt.Println(transactions)
	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	helper.ResponseHelper(w, err, transactions, 0)
}

func (h *handlerTransaction) GetTransactionByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	transactions, err := h.TransactionRepository.GetTransactionByUser(userId)
	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	helper.ResponseHelper(w, nil, transactions, 0)
}

func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusBadRequest)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", orderId)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			h.TransactionRepository.UpdateTransaction("success", orderId)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		h.TransactionRepository.UpdateTransaction("success", orderId)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", orderId)
	}

	helper.ResponseHelper(w, nil, notificationPayload, 0)
}
