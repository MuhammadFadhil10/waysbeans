package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(postgre.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/all-transactions", middleware.Auth(h.GetAllTransactions)).Methods("GET")
	r.HandleFunc("/transactions", middleware.Auth(h.GetTransactionByUser)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction-process", middleware.Auth(h.Notification)).Methods("POST")
}
