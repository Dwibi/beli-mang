CREATE TABLE IF NOT EXISTS order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(order_id),
    merchant_id INT NOT NULL REFERENCES merchants(id),
    item_id INT NOT NULL REFERENCES items(id),
    quantity INT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);
CREATE INDEX IF NOT EXISTS idx_order_items_merchant_id ON order_items(merchant_id);
CREATE INDEX IF NOT EXISTS idx_order_items_item_id ON order_items(item_id);