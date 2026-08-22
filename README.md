# 🏛️ Clean Architecture Task Manager API (Go + MongoDB)

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)](https://www.mongodb.com/)
[![JWT Auth](https://img.shields.io/badge/Auth-JWT_%2B_Bcrypt-black?style=for-the-badge&logo=jsonwebtokens&logoColor=white)](https://jwt.io/)
[![Mockery Testing](https://img.shields.io/badge/Tests-Mockery_%2B_Testify-brightgreen?style=for-the-badge)](https://github.com/vektra/mockery)

A robust, enterprise-grade RESTful Task Management API engineered in **Go (Golang)** adhering strictly to Uncle Bob's **Clean Architecture** and **Domain-Driven Design (DDD)** principles.

---

## 🎯 Key Architectural Pillars

- **Decoupled Layers**: Separation of concerns across **Delivery (HTTP/Controllers)**, **Usecases (Business Logic)**, **Domain (Entities/Interfaces)**, and **Repositories/Infrastructure (Database & Security)**.
- **Dependency Inversion**: High-level modules do not depend on low-level modules; both depend on abstract interfaces.
- **Test-Driven Design**: Full unit test coverage for controllers and usecases utilizing **Mockery** and **Testify** test suites without requiring a live database during CI.
- **Enterprise Security**: HMAC-SHA256 JWT access tokens, role-based authorization middleware (Admin/User), and bcrypt password hashing.

---

## 📐 System Architecture

```
task_manager_final/
├── Delivery/                  # Presentation Layer
│   ├── Controllers/           # HTTP Request Handlers & DTO validation
│   ├── routers/               # Gin Route Handlers & Middleware binding
│   └── main.go                # Application Entrypoint & Dependency Wiring
├── Domain/                    # Enterprise Business Rules & Entities
│   └── domain.go              # Entity definitions & Repository/Usecase Interfaces
├── Usecases/                  # Application Business Rules
│   ├── task_usecases.go       # Task business logic & validation
│   └── user_usecases.go       # User registration, login & token issuance
├── Repositories/              # Interface Adapters / Data Access Layer
│   ├── task_repository.go     # MongoDB Task persistence implementation
│   └── user_repository.go     # MongoDB User persistence implementation
├── Infrastructure/            # Frameworks & Drivers
│   ├── auth_middleware.go     # Gin JWT authentication middleware
│   ├── jwt_service.go         # Token generation & claim verification
│   └── password_service.go    # Bcrypt password hashing & comparison
└── Mocks/                     # Auto-generated Test Mocks (Mockery)
    ├── TaskRepositoryMock.go
    ├── TaskUsecaseMock.go
    ├── UserRepositoryMock.go
    └── UserUsecaseMock.go
```

---

## 🚀 REST API Endpoints

### 🔐 Authentication (`/auth`)
| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/v1/register` | Register new user account | No |
| `POST` | `/api/v1/login` | Authenticate & obtain JWT Bearer token | No |

### 📋 Task Management (`/tasks`)
| Method | Endpoint | Description | Role / Auth |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/v1/tasks` | Retrieve all user tasks (with pagination) | Bearer Token |
| `GET` | `/api/v1/tasks/:id` | Fetch detailed task by ID | Bearer Token |
| `POST` | `/api/v1/tasks` | Create a new task entity | Bearer Token |
| `PUT` | `/api/v1/tasks/:id` | Update task status, due date, description | Bearer Token |
| `DELETE`| `/api/v1/tasks/:id` | Delete task by ID | Admin / Owner |

---

## 🛠️ Quick Start & Local Setup

### Prerequisites
- **Go**: `1.22+`
- **MongoDB**: `6.0+` (or MongoDB Atlas URI)

### Installation
```bash
# Clone the repository
git clone https://github.com/hosama-adem/task_manager_final.git
cd task_manager_final

# Download dependencies
go mod download

# Set Environment Variables
export PORT=8080
export MONGO_URI="mongodb://localhost:27017"
export DB_NAME="task_manager_db"
export JWT_SECRET="your-super-secret-jwt-key"

# Run the API server
go run Delivery/main.go
```

---

## 🧪 Running Unit Tests

```bash
# Run all unit tests with verbose coverage
go test -v ./...

# Run controller and usecase tests specifically
go test -v ./Delivery/Controllers/...
go test -v ./Usecases/...
```

---

## 📄 License
Distributed under the MIT License. Created by [Hosama Adem](https://github.com/hosama-adem).
