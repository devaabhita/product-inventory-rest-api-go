package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon (Ping Gagal): ", err)
	}

	fmt.Println("Database Connected!")

	// membuat tabel if table not exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(10, 2) NOT NULL,
		stock INT NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Gagal menjalankan auto-migration tabel: ", err)
	}
}