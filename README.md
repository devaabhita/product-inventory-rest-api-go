# Product Inventory REST API

Backend API RESTful untuk manajemen inventaris produk yang dibangun menggunakan bahasa pemrograman **Go (Golang)** dan database **PostgreSQL**. Proyek ini dirancang menggunakan arsitektur yang bersih (_clean structure_) dan sepenuhnya diorkestrasi menggunakan **Docker & Docker Compose** untuk kemudahan _deployment_.

---

## Fitur Utama

- **CRUD Operations:** Mendukung manajemen data produk (List Products & Add Product).
- **Auto-Migration:** Tabel database `products` otomatis terbentuk saat aplikasi pertama kali dinyalakan.
- **Database Healthcheck:** Backend Go dikonfigurasi secara cerdas untuk menunggu database PostgreSQL benar-benar siap (_healthy_) sebelum menginisialisasi koneksi.
- **Multi-stage Docker Build:** Menggunakan teknik multi-stage build pada Dockerfile untuk menghasilkan _production image_ Alpine yang sangat ringan dan aman.

---

## Struktur Proyek

```text
product-inventory/
├── config/             # Konfigurasi aplikasi & database connection
│   └── database.go
├── controllers/        # Logika bisnis / Handler HTTP API
│   └── product.go
├── models/             # Definisi skema struktur data (Struct)
│   └── product.go
├── docker-compose.yml  # Orkestrasi kontainer aplikasi & database
├── dockerfile          # Instruksi blueprint Docker image backend
├── main.go             # Entrypoint utama aplikasi
└── .env                # Konfigurasi environment variable lokal
```

---

## Prasyarat (Prerequisites)

Sebelum menjalankan proyek ini, pastikan Anda sudah menginstal:

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/)
- [Postman](https://www.postman.com/) (untuk pengujian API)

---

## Cara Menjalankan Aplikasi

Anda tidak perlu menginstal Go atau PostgreSQL secara lokal di komputer Anda. Cukup gunakan Docker dengan langkah-langkah berikut:

1. **Clone Repositori Ini**

   ```bash
   git clone [https://github.com/devaabhita/product-inventory-rest-api-go.git](https://github.com/devaabhita/product-inventory-rest-api-go.git)
   cd product-inventory-rest-api-go
   ```

2. **Jalankan dengan Docker Compose**
   Eksekusi perintah berikut di terminal Anda:

   ```bash
   docker compose up --build
   ```

3. **Verifikasi Kontainer**
   Tunggu hingga terminal menampilkan log:
   ```text
   Database Connected!
   Server running on port 8080
   ```
   Aplikasi kini dapat diakses di `http://localhost:8080`.

---

## Panduan Pengujian API (Postman)

### 1. Mengambil Semua Produk (Get All Products)

- **Method:** `GET`
- **URL:** `http://localhost:8080/products`
- **Respon yang Diharapkan (Jika Kosong):** `[]`

### 2. Menambah Produk Baru (Create Product)

- **Method:** `POST`
- **URL:** `http://localhost:8080/products`
- **Headers:** `Content-Type: application/json`
- **Body (Raw JSON):**
  ```json
  {
    "name": "Laptop ASUS ROG",
    "price": 15000000,
    "stock": 10
  }
  ```

```

```
