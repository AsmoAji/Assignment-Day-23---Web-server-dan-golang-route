
# Inventory Management API - Golang

Proyek ini merupakan sistem backend untuk aplikasi manajemen inventaris. Dibangun dengan Golang menggunakan framework Gin dan GORM sebagai ORM, serta mendukung unggah dan unduh gambar produk.

---

## 🧱 Fitur

- CRUD Produk
- Manajemen Stok Inventaris
- Manajemen Pesanan
- Upload & Download Gambar Produk
- Validasi dan penyimpanan file (gambar)
- RESTful API
- Pengujian API dengan Postman

---

## 📂 Struktur Folder

```
inventory-app/
│
├── config/             # Konfigurasi database
├── controllers/        # Handler untuk API
├── models/             # Struktur data & model database
├── routes/             # Routing endpoint
├── uploads/            # Tempat penyimpanan gambar lokal
├── main.go             # Entry point aplikasi
├── go.mod              # Module Go
└── README.md           # Dokumentasi
```

---

## ⚙️ Setup Database (MySQL)

### 1. Buat Database Baru

Gunakan DBeaver atau CLI:

```sql
CREATE DATABASE inventory_db;
```

---

## 🛠️ Konfigurasi Project

### 2. Ubah Konfigurasi di `config/database.go`

```go
dsn := "root:@tcp(127.0.0.1:3306)/inventory_db?charset=utf8mb4&parseTime=True&loc=Local"
```

Jika kamu tidak menggunakan password, cukup kosongkan setelah `root:`. Jika pakai password, isi sesuai.

---

## 🚀 Menjalankan Aplikasi

1. Clone repo:
```bash
git clone https://github.com/username/inventory-app.git
cd inventory-app
```

2. Install dependensi:
```bash
go mod tidy
```

3. Jalankan aplikasi:
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

## 🔌 API Endpoint

### 📦 Produk

| Method | Endpoint            | Deskripsi                    |
|--------|---------------------|------------------------------|
| GET    | /products           | Ambil semua produk           |
| GET    | /products/:id       | Ambil produk berdasarkan ID  |
| POST   | /products           | Tambah produk                |
| PUT    | /products/:id       | Update produk                |
| DELETE | /products/:id       | Hapus produk                 |
| POST   | /products/:id/image | Upload gambar produk         |
| GET    | /products/:id/image | Download gambar produk       |

---

### 📊 Inventaris

| Method | Endpoint                  | Deskripsi                      |
|--------|---------------------------|--------------------------------|
| GET    | /inventory/:product_id    | Lihat stok untuk produk        |
| PUT    | /inventory/:product_id    | Update stok produk             |

---

### 🛒 Pesanan

| Method | Endpoint         | Deskripsi                     |
|--------|------------------|-------------------------------|
| POST   | /orders          | Tambah pesanan baru           |
| GET    | /orders/:id      | Ambil detail pesanan          |

---

## 🧪 Pengujian dengan Postman

1. Buka **Postman**
2. Gunakan endpoint sesuai tabel di atas
3. Untuk upload gambar:
   - Method: `POST`
   - URL: `http://localhost:8080/products/{id}/image`
   - Body: form-data → Key: `image`, Type: File

---

## 📸 Validasi Gambar

- Format yang didukung: `.jpg`, `.jpeg`, `.png`
- Maksimum ukuran file: 2MB

---

## 🧼 Error Handling

API akan mengembalikan pesan error dalam format JSON jika:
- File tidak dikirim
- Format file salah
- File terlalu besar
- Data tidak ditemukan

---

## 🧾 License

MIT License
