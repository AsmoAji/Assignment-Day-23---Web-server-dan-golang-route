# 📦 Inventory Management API

Sistem backend untuk aplikasi manajemen inventaris menggunakan Golang, Gin, dan MySQL. Menyediakan RESTful API untuk mengelola produk, inventaris, dan pesanan, serta mendukung upload gambar produk.

---

## 🚀 Fitur Utama

- CRUD Produk
- Manajemen Inventaris
- Pemesanan Produk
- Upload dan Download Gambar Produk
- API berbasis Gin Framework
- Database Relasional menggunakan GORM + MySQL

---

## 🧱 Struktur Proyek

inventory-app/
├── config/
├── controllers/
├── models/
├── routes/
├── uploads/
├── main.go
├── go.mod
└── README.md

---

## 🛠️ Cara Menjalankan Proyek

### 1. Clone Repository
git clone https://github.com/AsmoAji/Assignment-Day-23---Web-server-dan-golang-route.git
cd inventory-app

### 2. Instal Dependensi
go mod tidy

### 3. Setup MySQL
- Pastikan MySQL aktif di localhost:3306
- Buat database: inventory_db

CREATE DATABASE inventory_db;

### 4. Konfigurasi `config/database.go`
dsn := "root:@tcp(127.0.0.1:3306)/inventory_db?charset=utf8mb4&parseTime=True&loc=Local"

### 5. Jalankan Aplikasi
go run main.go

---

## 📬 API Endpoint

Produk:
- POST    /products
- GET     /products
- GET     /products/:id
- PUT     /products/:id
- DELETE  /products/:id

Inventaris:
- GET     /inventory/:product_id
- PUT     /inventory/:product_id

Pesanan:
- POST    /orders
- GET     /orders/:id

Gambar Produk:
- POST    /products/:id/image
- GET     /products/:id/image

---

## 🧪 Pengujian API

Gunakan Postman untuk menguji endpoint.

---

## 🧑‍💻 Author

Nama: _[Tri Putranda Asmo Aji]_
