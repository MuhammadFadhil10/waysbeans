package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/postgre"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProducts(postgre.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", h.GetProducts)
}
