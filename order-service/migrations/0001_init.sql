CREATE TABLE orders (
                        id VARCHAR(255) PRIMARY KEY,
                        idempotency_key VARCHAR(255) UNIQUE, 
                        customer_id VARCHAR(255) NOT NULL,
                        item_name VARCHAR(255) NOT NULL,
                        amount BIGINT NOT NULL,
                        status VARCHAR(50) NOT NULL,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);