package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"waysbeans/database"
	"waysbeans/pkg/postgre"
	"waysbeans/routes"

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

	PORT := os.Getenv("PORT")

	// route init
	routes.RoutesInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running on port ", PORT)

	http.ListenAndServe(":"+PORT, r)
}
