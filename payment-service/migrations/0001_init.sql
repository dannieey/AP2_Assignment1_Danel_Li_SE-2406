CREATE TABLE payments (
                          id VARCHAR(255) PRIMARY KEY,
                          order_id VARCHAR(255) NOT NULL,
                          transaction_id VARCHAR(255) NOT NULL,
                          amount BIGINT NOT NULL,
                          status VARCHAR(50) NOT NULL
);