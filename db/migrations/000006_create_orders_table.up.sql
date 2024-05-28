CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    calculated_estimate_id INT NOT NULL REFERENCES estimates(calculated_estimate_id),
    user_id INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);