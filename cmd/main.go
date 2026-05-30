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
	config.ConnectDB()
	db := config.DB
	defer db.Close()

	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/products", controllers.ListProducts(db)).Methods("GET")
	r.HandleFunc("/products", controllers.AddProduct(db)).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.GetProductByID(db)).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.UpdateProduct(db)).Methods("PUT")
	r.HandleFunc("/products/{id}", controllers.DeleteProduct(db)).Methods("DELETE")
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}