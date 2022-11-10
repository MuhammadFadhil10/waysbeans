package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"waysbeans/database"
	"waysbeans/pkg/postgre"
	"waysbeans/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	postgre.DatabaseInit()
	database.RunMigration()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	// cors
	var allowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})
	var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

	PORT := os.Getenv("PORT")

	// route init
	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running on port ", PORT)

	http.ListenAndServe(":"+PORT, handlers.CORS(allowedHeaders,allowedMethods,allowedOrigins)(r))
}
