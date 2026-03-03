# Personal Finance Tracker API

A RESTful backend service built with Go that allows users to track their income and expenses by category, generate monthly reports, and manage budgets.

---

## Tech Stack

- **Go** — Backend language
- **Gin** — HTTP web framework
- **PostgreSQL** — Relational database
- **GORM** — ORM for database operations
- **JWT** — Authentication
- **Swagger** — API documentation

---

## Features

- User registration and login with JWT authentication
- Each user gets 8 default categories (Food, Transport, Shopping, etc.) upon registration
- Full CRUD for categories and transactions
- Default categories are protected from deletion
- Monthly financial summary (total income, total expense, net balance)
- Category-based expense analysis
- Month-over-month comparison reports
- Budget limit management per category

---

## Project Structure

```
finance-tracker/
├── cmd/
│   └── main.go               # Entry point, route definitions
├── config/
│   └── config.go             # Database connection
├── docs/                     # Auto-generated Swagger docs
├── internal/
│   ├── handlers/             # HTTP layer (request/response)
│   │   ├── auth_handler.go
│   │   ├── category_handler.go
│   │   ├── transaction_handler.go
│   │   └── report_handler.go
│   ├── middleware/
│   │   └── auth_middleware.go  # JWT validation
│   ├── models/
│   │   ├── user.go             # Database models
│   │   ├── budget.go
│   │   └── dto.go              # Data Transfer Objects
│   └── services/               # Business logic layer
│       ├── user_service.go
│       ├── category_service.go
│       ├── transaction_service.go
│       └── report_service.go
├── .env
└── go.mod
```

---

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL

### Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/finance-tracker.git
cd finance-tracker
```

2. Create a `.env` file:
```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=finance_tracker
DB_PORT=5432
JWT_SECRET=your_secret_key
PORT=8080
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
go run cmd/main.go
```

5. Visit Swagger UI:
```
http://localhost:8080/swagger/index.html
```

---

## API Endpoints

### Auth
| Method | URL | Description |
|--------|-----|-------------|
| POST | /register | Register a new user |
| POST | /login | Login and receive JWT token |

### Categories
| Method | URL | Description |
|--------|-----|-------------|
| GET | /categories | List all categories |
| POST | /categories | Create a category |
| PUT | /categories/:id | Update a category |
| DELETE | /categories/:id | Delete a category (default categories protected) |

### Transactions
| Method | URL | Description |
|--------|-----|-------------|
| GET | /transactions | List all transactions |
| POST | /transactions | Create a transaction |
| PUT | /transactions/:id | Update a transaction |
| DELETE | /transactions/:id | Delete a transaction |

### Reports
| Method | URL | Description |
|--------|-----|-------------|
| GET | /report/summary | Monthly income/expense summary |
| GET | /report/categories | Expense breakdown by category |
| GET | /report/comparison | Compare current vs previous month |
| POST | /report/budget | Set a budget limit for a category |
| GET | /report/budget | Get budget status for the month |

---

## Authentication

All endpoints except `/register` and `/login` require a Bearer token in the Authorization header:

```
Authorization: Bearer <your_token>
```

---

## Architecture

The project follows a layered architecture:

```
Request → Handler → Service → Database
```

- **Handler** — Handles HTTP request/response, reads input, sends output
- **Service** — Contains business logic, interacts with the database
- **Middleware** — Validates JWT token on protected routes

---

## Security

- Passwords are hashed using **bcrypt** before storing
- JWT tokens expire after **24 hours**
- All database queries filter by `user_id` to prevent unauthorized access (IDOR protection)
- Token payload contains only `user_id` and `exp` — no sensitive data

---

## API Documentation

Swagger UI is available at:
```
http://localhost:8080/swagger/index.html
```

To regenerate docs after changes:
```bash
swag init -g main.go -d cmd,internal/handlers,internal/models
```
