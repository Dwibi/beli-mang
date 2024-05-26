DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'merchant_category') THEN
        CREATE TYPE merchant_category AS ENUM (
            'SmallRestaurant',
            'MediumRestaurant',
            'LargeRestaurant',
            'MerchandiseRestaurant',
            'BoothKiosk',
            'ConvenienceStore'
        );
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS merchants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    merchant_category merchant_category NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    location_lat FLOAT8 NOT NULL,
    location_long FLOAT8 NOT NULL
);