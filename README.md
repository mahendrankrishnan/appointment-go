# 🩺 Go Appointment API

A lightweight REST API built in Go for managing appointments. The API supports full CRUD operations and includes Swagger documentation for testing and development.

---

## 🚀 Features

- Create, update, delete, and retrieve appointments
- RESTful routes using a Go router
- Swagger UI for interactive API docs

---

## 📦 Endpoints

| Method | Endpoint         | Description                      |
|--------|------------------|----------------------------------|
| `GET`  | `/swagger/*`     | Swagger UI for API docs          |
| `POST` | `/`              | Create a new appointment         |
| `PUT`  | `/{id}`          | Update an appointment by ID      |
| `GET`  | `/`              | Retrieve all appointments        |
| `GET`  | `/{id}`          | Retrieve a specific appointment  |
| `DELETE` | `/{id}`        | Delete an appointment by ID      |

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Router**: Chi or similar Go router
- **Swagger**: `http-swagger` for API documentation

---

## 🧪 Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/mahendrankrishnan/go-appointment-api.git
cd go-appointment-api
