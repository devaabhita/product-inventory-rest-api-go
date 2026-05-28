package models

import "database/sql"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// GetProducts mengambil semua data
func GetProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// CreateProduct menambah data baru
func CreateProduct(db *sql.DB, p Product) error {
	_, err := db.Exec("INSERT INTO products (name, price, stock) VALUES ($1, $2, $3)", 
		p.Name, p.Price, p.Stock)
	return err
}