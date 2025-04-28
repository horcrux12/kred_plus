
# 📦 Kred Plus

## 🚀 Overview
**Kred Plus** adalah aplikasi pembiayaan white goods, motor, dan mobil berbasis Go (Golang) yang menggunakan database MySQL.  
Project ini siap dijalankan menggunakan **Docker** dan **Docker Compose**.

---

## 📂 Struktur Directory

```plaintext
.
├── app/            # Business logic
├── config/         # Configuration (env loader, database connection)
├── db/             # Database migrations (if any)
├── dto/            # Data Transfer Objects
├── handler/        # HTTP handlers
├── lib/            # Library utilities
├── model/          # ORM models
├── repository/     # Data layer / repository pattern
├── router/         # HTTP routing setup
├── service/        # Service layer
├── uploads/        # Upload directory (eg. for customer KTP/selfie)
├── main.go         # App entrypoint
├── go.mod
├── go.sum
├── docker-compose.yml
└── Dockerfile
```

---

## 🐳 How to Run with Docker

### 1. Build Docker Image

```bash
docker build -t kred_plus .
```

### 2. Start Using Docker Compose

```bash
docker-compose up --build
```

Docker Compose akan menjalankan:
- **App** di port `8910:8080`
- **MySQL** di port `3309:3306`

### 3. Access App

```bash
http://localhost:8080
```

---

# 📖 API Contract

### 🔹 POST `/customers`
**Create new customer**

Request:

```json
{
  "nik": "string",
  "full_name": "string",
  "legal_name": "string",
  "place_of_birth": "string",
  "date_of_birth": "YYYY-MM-DD",
  "salary": "number",
  "ktp_photo": "file upload",
  "selfie_photo": "file upload"
}
```

Response:

```json
{
  "message": "Customer created successfully",
  "data": {
    "id": "uuid",
    "full_name": "string"
  }
}
```

---

### 🔹 POST `/transactions`
**Create a new transaction (loan request)**

Request:

```json
{
  "customer_id": "uuid",
  "otr_price": 2000000,
  "admin_fee": 50000,
  "asset_name": "Motor Honda Beat",
  "tenor_months": 4
}
```

Response:

```json
{
  "message": "Transaction created successfully",
  "data": {
    "contract_number": "TXN-123456",
    "monthly_installment": 550000,
    "interest_amount": 50000
  }
}
```

---

### 🔹 GET `/customers/{id}`
**Retrieve customer data**

Response:

```json
{
  "id": "uuid",
  "full_name": "string",
  "nik": "string",
  "salary": "number",
  "limits": {
    "1_month": 100000,
    "2_months": 200000,
    "3_months": 300000,
    "4_months": 400000
  }
}
```

---

# 🛠 Requirements

- Go 1.23+
- Docker & Docker Compose
- MySQL 5.7+
- Make sure `.env` file diisi sesuai kebutuhan
