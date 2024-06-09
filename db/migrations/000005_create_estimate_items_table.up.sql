CREATE TABLE IF NOT EXISTS estimate_items (
    estimate_item_id SERIAL PRIMARY KEY,
    calculated_estimate_id INT NOT NULL REFERENCES estimates(calculated_estimate_id),
    merchant_id INT NOT NULL REFERENCES merchants(id),
    item_id INT NOT NULL REFERENCES items(id),
    quantity INT NOT NULL,
    is_starting_point BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_estimate_items_calculated_estimate_id ON estimate_items(calculated_estimate_id);
CREATE INDEX IF NOT EXISTS idx_estimate_items_merchant_id ON estimate_items(merchant_id);
CREATE INDEX IF NOT EXISTS idx_estimate_items_item_id ON estimate_items(item_id);