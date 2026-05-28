package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"product-inventory/config"
	"product-inventory/controllers"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r := mux.NewRouter()

	// Routing
	r.HandleFunc("/products", controllers.ListProducts(db)).Methods("GET")
	r.HandleFunc("/products", controllers.AddProduct(db)).Methods("POST")

	port := os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}