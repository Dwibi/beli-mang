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
    price INTEGER NOT NULL 
);