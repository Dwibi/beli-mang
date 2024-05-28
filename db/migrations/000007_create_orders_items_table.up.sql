CREATE TABLE IF NOT EXISTS order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders(order_id),
    merchant_id INT NOT NULL REFERENCES merchants(id),
    item_id INT NOT NULL REFERENCES items(id),
    quantity INT NOT NULL
);