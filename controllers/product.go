package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product-inventory/models"
	"strconv"

	"github.com/gorilla/mux"
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

		if products == nil {
			products = []models.Product{}
		}

		json.NewEncoder(w).Encode(products)
	}
}

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

func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID tidak valid", http.StatusBadRequest)
			return
		}

		var p models.Product
		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Payload JSON tidak valid", http.StatusBadRequest)
			return
		}

		query := "UPDATE products SET name=$1, price=$2, stock=$3 WHERE id=$4"
		result, err := db.Exec(query, p.Name, p.Price, p.Stock, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
			return
		}

		p.ID = id
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Produk berhasil diperbarui",
			"data":    p,
		})
	}
}

// menghapus produk berdasarkan ID
func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID tidak valid", http.StatusBadRequest)
			return
		}

		query := "DELETE FROM products WHERE id=$1"
		result, err := db.Exec(query, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Produk berhasil dihapus",
		})
	}
}