package orderitemrepository

/*
CREATE TABLE IF NOT EXISTS order_items (

	order_item_id SERIAL PRIMARY KEY,
	order_id INT NOT NULL REFERENCES orders(order_id),
	merchant_id INT NOT NULL REFERENCES merchants(id),
	item_id INT NOT NULL REFERENCES items(id),
	quantity INT NOT NULL

);

CREATE TABLE IF NOT EXISTS estimate_items (

	estimate_item_id SERIAL PRIMARY KEY,
	calculated_estimate_id INT NOT NULL REFERENCES estimates(calculated_estimate_id),
	merchant_id INT NOT NULL REFERENCES merchants(id),
	item_id INT NOT NULL REFERENCES items(id),
	quantity INT NOT NULL,
	is_starting_point BOOLEAN NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);
*/
func (i sOrderItemRepository) Create(orderId int, CalculatedEstimateId string) error {
	query := `INSERT INTO order_items (order_id, merchant_id, item_id, quantity)
	SELECT
		$1 AS order_id,
		merchant_id,
		item_id,
		quantity
	FROM estimate_items
	WHERE estimate_items.calculated_estimate_id = $2;`

	_, err := i.DB.Exec(query, orderId, CalculatedEstimateId)

	if err != nil {
		return err
	}

	return nil
}
