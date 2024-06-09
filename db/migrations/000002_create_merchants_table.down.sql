DROP TABLE IF EXISTS merchants CASCADE;
DROP TYPE IF EXISTS merchant_category CASCADE;

DROP INDEX IF EXISTS idx_merchants_location_lat_long;
DROP INDEX IF EXISTS idx_merchants_name;
DROP INDEX IF EXISTS idx_merchants_merchant_category;