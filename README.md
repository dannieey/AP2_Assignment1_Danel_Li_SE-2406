# AP2 Assignment 1: Microservices Order & Payment System

**Student:** Li Danel  
**Group:** SE-2406  
**Instructor:** Taubakabyl Nurlybek  
 

## 1. Project Overview
This project is a distributed system consisting of two independent microservices: **Order Service** and **Payment Service**. It is built following **Clean Architecture** principles and demonstrates synchronous REST communication with strict data isolation.

## 2. Key Features
- **Clean Architecture:** Strict separation of Domain, Usecase, Repository, and Transport layers.
- **Database per Service:** Independent PostgreSQL instances for each service to ensure data ownership.
- **Resilient Communication:** Implementation of custom HTTP clients with a **2-second timeout**.
- **Failure Handling:** Returns `503 Service Unavailable` if the Payment Service is unreachable or times out.
- **Bonus Feature (Idempotency):** Integrated `X-Idempotency-Key` header to prevent duplicate orders and payments.

## 3. Architecture & Bounded Contexts
- **Order Context:** Manages order placement and status (`Pending`, `Paid`, `Failed`, `Cancelled`).
- **Payment Context:** Authorizes transactions and enforces a limit of 100,000 units (1,000.00).
- **No Shared Code:** Each service maintains its own domain entities and logic, avoiding a "distributed monolith" design.

---

## 4. How to Run (Infrastructure)

**Note:** You do NOT need to install PostgreSQL on your machine. This project uses **Docker Desktop** to manage databases.

### Step 1: Start Databases via Docker
Run the following command in the root folder:
```bash
docker-compose up -d
```
* **Order DB:** `localhost:5431` (User: user, Pass: 0000)
* **Payment DB:** `localhost:5432` (User: user, Pass: 0000)

### Step 2: Run Microservices
Open two separate terminals:

**Terminal 1 (Payment Service):**
```bash
cd payment-service
go run cmd/main.go
```

**Terminal 2 (Order Service):**
```bash
cd order-service
go run cmd/main.go
```

