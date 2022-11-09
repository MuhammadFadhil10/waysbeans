package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(postgre.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", h.AddToCart).Methods("POST")
	r.HandleFunc("/carts", h.GetCarts).Methods("GET")
	r.HandleFunc("/cart/{cartId}", h.UpdateCartQty).Methods("PATCH")
}
