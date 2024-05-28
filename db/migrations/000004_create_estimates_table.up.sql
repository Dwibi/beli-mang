CREATE TABLE IF NOT EXISTS estimates (
    calculated_estimate_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    total_price DECIMAL(10, 3) NOT NULL DEFAULT 0.0,
    estimated_delivery_time INT NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);