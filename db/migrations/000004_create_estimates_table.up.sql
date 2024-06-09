CREATE TABLE IF NOT EXISTS estimates (
    calculated_estimate_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    total_price NUMERIC NOT NULL,
    estimated_delivery_time NUMERIC NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_estimates_user_id ON estimates(user_id);
