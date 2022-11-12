package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(postgre.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", middleware.Auth(h.AddToCart)).Methods("POST")
	r.HandleFunc("/carts", middleware.Auth(h.GetCarts)).Methods("GET")
	r.HandleFunc("/cart", middleware.Auth(h.GetCartByUser)).Methods("GET")
	r.HandleFunc("/cart/{cartId}", middleware.Auth(h.UpdateCartQty)).Methods("PATCH")
	r.HandleFunc("/cart/delete/{cartId}", middleware.Auth(h.DeleteCartByID)).Methods("DELETE")
	r.HandleFunc("/cart/clear/{userId}", middleware.Auth(h.DeleteCartByUser)).Methods("DELETE")
}
