package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(postgre.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/auth/register", h.Register).Methods("POST")
	r.HandleFunc("/auth/login", h.Login).Methods("POST")
}
