package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(postgre.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")
}