# Product API — Go Learning Project

A REST API built in Go for managing products. Built as part of a structured Go API learning roadmap.

---

## Requirements to rewrite this from scratch

### Structs & Data
- Create a `Product` struct with 4 fields: ID (int), Name (string), Price (float64), Quantity (int) — all with JSON tags
- Create a global `products` variable as an empty slice of Product

### Helper Functions
- `findProductByID(id int)` — returns index (int) and pointer to Product. Loop through products, return index and address of product if found, return -1 and nil if not found
- `getIDFromURL(r)` — extract the path after `/api/v1/products/`, convert to int, return int and error

### `allProductsHandler`
- Set Content-Type header
- GET → encode and return all products
- POST → decode body into new product, handle error, assign ID as len+1, append to slice, return 201 with created product
- default → 405

### `singleProductHandler`
- Set Content-Type header
- Extract and validate ID from URL, return 400 if invalid
- Find product, return 404 if not found
- GET → return the product
- PUT → decode body, handle error, keep same ID, update slice, return updated product
- DELETE → remove from slice using index, return 204
- default → 405

### `RequestHandler`
- Register `/api/v1/products` → allProductsHandler
- Register `/api/v1/products/` → singleProductHandler
- Start server on port 8083 with log.Fatal

### `main`
- Call RequestHandler

### Imports needed
`encoding/json`, `log`, `net/http`, `strconv`

---

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/products` | Get all products |
| POST | `/api/v1/products` | Create a new product |
| GET | `/api/v1/products/{id}` | Get a single product |
| PUT | `/api/v1/products/{id}` | Update a product |
| DELETE | `/api/v1/products/{id}` | Delete a product |

---

## Status Codes Used

| Code | Constant | When |
|------|----------|------|
| 200 | `http.StatusOK` | Default success |
| 201 | `http.StatusCreated` | Product created |
| 204 | `http.StatusNoContent` | Product deleted |
| 400 | `http.StatusBadRequest` | Invalid ID or bad body |
| 404 | `http.StatusNotFound` | Product not found |
| 405 | `http.StatusMethodNotAllowed` | Wrong HTTP method |

---

## How to Run

```bash
go run main.go
```

Server starts on port `8083`.

---

## How to Test (Windows PowerShell)

```powershell
# Create a product
Invoke-WebRequest -Uri http://localhost:8083/api/v1/products -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"name":"Laptop","price":999.99,"quantity":5}' -UseBasicParsing

# Get all products
Invoke-WebRequest -Uri http://localhost:8083/api/v1/products -Method GET -UseBasicParsing

# Get one product
Invoke-WebRequest -Uri http://localhost:8083/api/v1/products/1 -Method GET -UseBasicParsing

# Update a product
Invoke-WebRequest -Uri http://localhost:8083/api/v1/products/1 -Method PUT -Headers @{"Content-Type"="application/json"} -Body '{"name":"Gaming Laptop","price":1299.99,"quantity":3}' -UseBasicParsing

# Delete a product
Invoke-WebRequest -Uri http://localhost:8083/api/v1/products/2 -Method DELETE -UseBasicParsing
```

---

## Learning Roadmap

| Module | Topic | Status |
|--------|-------|--------|
| 1 | HTTP & CRUD Fundamentals | ✅ Done |
| 2 | Better Routing (chi) | ⏳ Next |
| 3 | Database (PostgreSQL + GORM) | 🔒 Upcoming |
| 4 | Project Structure | 🔒 Upcoming |
| 5 | Auth (JWT) | 🔒 Upcoming |
| 6 | Production Ready (Docker) | 🔒 Upcoming |

---

## Key Concepts Learned in Module 1

- HTTP server setup with `net/http`
- Structs and JSON tags
- In-memory CRUD using a slice
- Request and response handling
- URL path parsing
- HTTP status codes
- Pointer vs value types (`*Product` vs `Product`)
- Why headers must be set before the response body