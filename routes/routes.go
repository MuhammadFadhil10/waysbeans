package routes

import "github.com/gorilla/mux"

func RoutesInit(r *mux.Router) {
	ProductRoutes(r)
	AuthRoutes(r)
	CartRoutes(r)
	TransactionRoutes(r)

}
