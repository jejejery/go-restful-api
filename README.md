# Product Management API

## 📌 Deskripsi
API berbasis Golang untuk mengelola produk menggunakan **Fiber**, **GORM**, dan **MySQL**. API ini mendukung operasi CRUD (Create, Read, Update, Delete) dengan arsitektur **MVC** serta menerapkan **Repository Pattern**.

---

## 🚀 Teknologi yang Digunakan
- [Golang](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)

---

## 🚀Get Dependency
```shell

go get -u github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 🚀Update Dependency
```shell

go mod tidy
```

## 📂 Struktur Direktori
```
project-root/
│── controllers/
│   └── product_controller.go
│── database/
│   └── database.go
│── models/
│   └── product.go
│── repositories/
│   └── product_repository.go
│── routes/
│   └── routes.go
│── services/
│   └── product_service.go
│── main.go
│── go.mod
│── go.sum
│── README.md
```

---

## 🛠️ Instalasi & Menjalankan Proyek

### 1️⃣ Clone Repository
```sh
git clone https://github.com/username/product-management-api.git
cd product-management-api
```

### 2️⃣ Install Dependency
```sh
go mod tidy
```

### 3️⃣ Konfigurasi Database
Edit file `database/database.go`, ubah `dsn` sesuai dengan kredensial MySQL Anda:
```go
dsn := "user:password@tcp(localhost:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
```

### 4️⃣ Jalankan Aplikasi
```sh
go run main.go
```

API akan berjalan di: `http://localhost:8080`

---

## 🔥 Endpoint API
| Metode | Endpoint       | Deskripsi               |
|--------|--------------|-------------------------|
| POST   | `/products/` | Tambah produk baru      |
| GET    | `/products/` | Ambil semua produk      |
| GET    | `/products/:id` | Ambil produk berdasarkan ID |
| PUT    | `/products/:id` | Update produk berdasarkan ID |
| DELETE | `/products/:id` | Hapus produk berdasarkan ID |

### 📌 Contoh Request
#### 🔹 Tambah Produk Baru
**Request:**
```json
POST /products/
{
  "product_id": "P001",
  "name": "Laptop Gaming",
  "description": "Laptop dengan spesifikasi tinggi",
  "price": 15000000,
  "stock_qty": 10,
  "category": "Electronics",
  "sku": "LAP123",
  "tax_rate": 10.0
}
```

**Response:**
```json
{
  "product_id": "P001",
  "name": "Laptop Gaming",
  "description": "Laptop dengan spesifikasi tinggi",
  "price": 15000000,
  "stock_qty": 10,
  "category": "Electronics",
  "sku": "LAP123",
  "tax_rate": 10.0
}
```

---

## ✨ Kontributor
- **Ahmad Roni Purwanto** - Full Stack Developer

---

## 📜 Lisensi
MIT License. Silakan gunakan dan modifikasi proyek ini sesuai kebutuhan Anda.

