package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(postgre.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/profile", middleware.Auth(h.GetProfile)).Methods("GET")
}
