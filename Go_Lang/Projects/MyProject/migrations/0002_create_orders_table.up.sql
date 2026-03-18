CREATE TABLE orders(
    id VARCHAR(50) PRIMARY KEY,
    facility_code VARCHAR(50),
    status VARCHAR(50),
    created_at TIMESTAMP,
    FOREIGN KEY (facility_code) REFERENCES facilities(code) ON DELETE SET NULL
)