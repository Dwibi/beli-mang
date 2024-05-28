package estimateitemsrepository

func (i sEstimateItemsRepository) FindTotalPrice(calculatedEstimateId int) (float64, error) {
	query := "SELECT SUM(ei.quantity * i.price) AS total_price FROM estimate_items ei JOIN items i ON ei.item_id = i.id WHERE ei.calculated_estimate_id = $1;"

	var totalPrice float64
	err := i.DB.QueryRow(query, calculatedEstimateId).Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil
}
