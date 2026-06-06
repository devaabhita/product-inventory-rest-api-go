# Product Inventory REST API

Backend API RESTful untuk manajemen inventaris produk yang dibangun menggunakan bahasa pemrograman **Go (Golang)** dan database **PostgreSQL**. Proyek ini dirancang menggunakan arsitektur yang bersih (*clean structure*) dan sepenuhnya diorkestrasi menggunakan **Docker & Docker Compose** untuk kemudahan *deployment*.

---

## Fitur Utama

- **CRUD Operations:** Mendukung manajemen data produk secara lengkap (Create, Read, Update, Delete).
- **Auto-Migration:** Tabel database `products` otomatis terbentuk saat aplikasi pertama kali dinyalakan.
- **Database Healthcheck:** Backend Go dikonfigurasi secara cerdas untuk menunggu database PostgreSQL benar-benar siap (*healthy*) sebelum menginisialisasi koneksi.
- **Multi-stage Docker Build:** Menggunakan teknik multi-stage build pada Dockerfile untuk menghasilkan *production image* Alpine yang sangat ringan dan aman.

---

## Project Structure

Berikut adalah struktur direktori dan berkas dari proyek **PRODUCT-INVENTORY**:

```text
product-inventory/
├── cmd/
│   └── main.go              # Entrypoint utama aplikasi (menjalankan server)
├── config/
│   └── database.go          # Konfigurasi aplikasi & koneksi database
├── controllers/
│   └── product.go           # Logika bisnis / Handler HTTP API untuk produk
├── models/
│   └── product.go           # Definisi skema struktur data (Struct) produk
├── .dockerignore            # Berkas dan folder yang diabaikan oleh Docker
├── .env                     # Konfigurasi environment variable lokal
├── .gitignore               # Berkas dan folder yang diabaikan oleh Git
├── docker-compose.yml       # Orkestrasi kontainer aplikasi & database
├── dockerfile               # Instruksi blueprint Docker image backend
├── go.mod                   # Berkas manajemen modul dan dependensi Go
└── go.sum                   # Berkas checksum untuk verifikasi dependensi Go
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

### 2. Mengambil Satu Produk Berdasarkan ID (Get Product By ID)

- **Method:** `GET`
- **URL:** `http://localhost:8080/products/1` *(Ganti angka 1 dengan ID produk yang ingin dicari)*
- **Respon jika data ditemukan:**
  ```json
  {
    "id": 1,
    "name": "Laptop ASUS ROG",
    "price": 15000000,
    "stock": 10,
    "created_at": "2026-05-30T19:00:00Z"
  }
  ```
- **Respon jika data tidak ada:**
  ```text
  Produk tidak ditemukan
  ```

### 3. Menambah Produk Baru (Create Product)

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

### 4. Memperbarui Data Produk (Update Product)

- **Method:** `PUT`
- **URL:** `http://localhost:8080/products/1` *(Ganti angka 1 dengan ID produk yang ingin diubah)*
- **Headers:** `Content-Type: application/json`
- **Body (Raw JSON):**
  ```json
  {
    "name": "Laptop ASUS TUF",
    "price": 16000000,
    "stock": 8
  }
  ```

### 5. Menghapus Produk (Delete Product)

- **Method:** `DELETE`
- **URL:** `http://localhost:8080/products/1` *(Ganti angka 1 dengan ID produk yang ingin dihapus)*
- **Respon yang Diharapkan:**
  ```json
  {
    "message": "Produk berhasil dihapus"
  }
  ```
