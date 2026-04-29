# Livecode 2 - Tourism Booking REST API

Go + Echo + GORM + PostgreSQL + JWT

## Struktur Project

```
livecode2/
├── main.go                         # Entry point, routing, JWT middleware
├── config/
│   └── database.go                 # Koneksi database PostgreSQL
├── domain/                         # Interface contracts (Repository, UseCase, Handler)
│   ├── auth.go
│   ├── booking.go
│   ├── tour.go
│   └── report.go
├── model/
│   ├── users/
│   │   ├── entity_model.go         # Struct User & Customer (GORM)
│   │   └── http_model.go           # Request/Response struct
│   ├── bookings/
│   │   ├── entity_model.go         # Struct Booking, TourBooking, Payment
│   │   └── http_model.go
│   └── tours/
│       ├── entity_model.go         # Struct Tour
│       └── http_model.go
├── repository/db/                  # Data layer (SQL/GORM queries)
│   ├── auth_repository.go
│   ├── booking_repository.go
│   ├── tour_repository.go
│   └── report_repository.go
├── usecase/                        # Business logic layer
│   ├── auth_usecase.go
│   ├── booking_usecase.go
│   ├── tour_usecase.go
│   └── report_usecase.go
└── delivery/http/                  # HTTP Handler layer
    ├── auth_handler.go
    ├── booking_handler.go
    ├── tour_handler.go
    └── report_handler.go
```

## Setup & Run

```bash
# 1. Install dependencies
go mod tidy

# 2. Set environment variable
export DATABASE_URL="postgres://ud1b5uav6a245k:PASSWORD@host:5432/dbname?sslmode=require"

# 3. Jalankan server
go run main.go
# Server berjalan di http://localhost:8080
```

## Endpoints

| Method | Path | Auth | Deskripsi |
|--------|------|------|-----------|
| POST | `/users/register` | ❌ | Register user baru |
| POST | `/users/login` | ❌ | Login, mendapatkan JWT token |
| GET | `/bookings` | ✅ JWT | Semua booking milik user login |
| GET | `/bookings/unpaid` | ✅ JWT | Booking belum dibayar milik user login |
| GET | `/tours/earning` | ✅ JWT | Pendapatan tiap tour |
| GET | `/reports/total-customers` | ✅ JWT | Total pelanggan terdaftar |
| GET | `/reports/bookings-per-tour` | ✅ JWT | Total booking per tour |

## Cara Test

```bash
# 1. Register
curl -X POST http://localhost:8080/users/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@test.com","password":"pass123","phone_number":"08123456789","address":"Jakarta"}'

# 2. Login (gunakan akun yang baru diregister)
TOKEN=$(curl -s -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@test.com","password":"pass123"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['token'])")

# 3. Get All Bookings
curl http://localhost:8080/bookings \
  -H "Authorization: Bearer $TOKEN"

# 4. Get Unpaid Bookings
curl http://localhost:8080/bookings/unpaid \
  -H "Authorization: Bearer $TOKEN"

# 5. Tour Earnings
curl http://localhost:8080/tours/earning \
  -H "Authorization: Bearer $TOKEN"

# 6. Total Customers
curl http://localhost:8080/reports/total-customers \
  -H "Authorization: Bearer $TOKEN"

# 7. Bookings Per Tour
curl http://localhost:8080/reports/bookings-per-tour \
  -H "Authorization: Bearer $TOKEN"
```

## Error Codes

| Code | Arti |
|------|------|
| 400 | Bad Request (input tidak valid / email sudah terdaftar) |
| 401 | Unauthorized (JWT tidak ada / tidak valid) |
| 404 | Not Found (data tidak ditemukan) |
| 500 | Internal Server Error |
