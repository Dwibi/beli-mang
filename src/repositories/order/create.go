package orderrepository

func (i sOrderRepository) Create(reqUserId int, CalculatedEstimateId string) (int, error) {
	query := `INSERT INTO orders (user_id, calculated_estimate_id) VALUES ($1, $2) RETURNING order_id`

	var orderId int

	err := i.DB.QueryRow(query, reqUserId, CalculatedEstimateId).Scan(&orderId)

	if err != nil {
		return 0, err
	}

	return orderId, nil
}
