package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product-inventory/models"
)

func ListProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := db.Query("SELECT id, name, price, stock, created_at FROM products ORDER BY id DESC")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var p models.Product
			err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			products = append(products, p)
		}

		//jika data masih kosong, kembalikan array kosong [] bukan null
		if products == nil {
			products = []models.Product{}
		}

		json.NewEncoder(w).Encode(products)
	}
}

//memasukkan data produk baru ke database
func AddProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p models.Product
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Payload JSON tidak valid", http.StatusBadRequest)
			return
		}

		query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id, created_at"
		err = db.QueryRow(query, p.Name, p.Price, p.Stock).Scan(&p.ID, &p.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	}
}