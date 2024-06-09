DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_category') THEN
        CREATE TYPE product_category AS ENUM (
            'Beverage',
            'Food',
            'Snack',
            'Condiments',
            'Additions'
        );
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    product_category product_category NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    merchant_id INT NOT NULL REFERENCES merchants(id)
);

CREATE INDEX IF NOT EXISTS idx_items_name ON items(name);
CREATE INDEX IF NOT EXISTS idx_items_product_category ON items(product_category);
CREATE INDEX IF NOT EXISTS idx_items_merchant_id ON items(merchant_id);