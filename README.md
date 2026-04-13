# AP2 Assignment 2: gRPC Migration & Order Tracking System

**Student:** Li Danel  
**Group:** SE-2406  
**Instructor:** Taubakabyl Nurlybek  

---

## 1. Project Overview
This project is an evolution of the Microservices Order & Payment System. In this assignment, the system was migrated from REST-to-REST communication to a hybrid **REST + gRPC** architecture, following the **Contract-First** principle and implementing **Server-side Streaming**.

## 2. New Key Features (Assignment 2)
- **Contract-First Flow:** Protos are managed in a dedicated repository. Code is generated and imported as a Go module.
- **gRPC Unary Communication:** Order Service now calls Payment Service via gRPC for high-performance, type-safe authorization.
- **Server-side Streaming:** Real-time order tracking implemented. Clients can subscribe to order status updates via a persistent gRPC stream.
- **Event-Driven Streaming:** Updates are tied to actual database changes using Go Channels (no fake sleep loops).
- **Centralized Configuration:** All gRPC addresses, ports, and DB credentials are managed via `.env` files.

## 3. Architecture & Technologies
- **Communication:** - **External:** REST (Gin) for public API endpoints.
  - **Internal:** gRPC (Protobuf) for inter-service calls.
  - **Streaming:** gRPC Server Streams for real-time notifications.
- **Contracts:** Protos hosted at `github.com/dannieey/assignment2-generated`.
- **Concurrency:** Goroutines and Channels for managing stream subscribers in the Order Service.

---

## 4. How to Run (Infrastructure)

### Step 1: Start Databases via Docker
```bash
docker-compose up -d
```

### Step 2: Set up Environment Variables (`.env`)
Make sure your `order-service/.env` contains:
```env
PAYMENT_SERVICE_ADDR=localhost:50052
ORDER_GRPC_PORT=:50051
HTTP_PORT=:8080
```

### Step 3: Run Microservices

**Terminal 1 (Payment Service - gRPC Server):**
```bash
cd payment-service
go run cmd/main.go
```

**Terminal 2 (Order Service - gRPC Client & Server):**
```bash
cd order-service
go run cmd/main.go
```

**Terminal 3 (gRPC Tracking Client):**
```bash
cd order-service
go run cmd/client/main.go
```

---

## 5. API Testing (Postman)

### Create Order (Triggers gRPC Call & Stream)
- **Method:** `POST`
- **URL:** `http://localhost:8080/orders`
- **Headers:** `X-Idempotency-Key: unique-key-123`
- **Body:**
```json
{
    "customer_id": "danel_dev",
    "item_name": "Strawberry Bingsu",
    "amount": 2500
}
```

### Expected Results:
1. **Postman:** Receives `Status: Paid` (Confirmed via gRPC Unary call).
2. **Tracking Client:** Terminal 3 displays: `UPDATE RECEIVED: Order ID: ..., Status: Paid`.

---

## 6. Project Structure
- `order-service/internal/transport/grpc`: Implementation of the streaming server.
- `order-service/internal/client`: gRPC client for Payment Service.
- `order-service/cmd/client`: CLI tool to demonstrate real-time tracking.
```

