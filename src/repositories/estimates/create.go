package estimatesrepository

func (i sEstimatesRepository) Create(userId int, estimatedDeliveryTime int) (int, error) {
	var estimateId int
	err := i.DB.QueryRow("INSERT INTO estimates (user_id, estimated_delivery_time) VALUES ($1, $2) RETURNING calculated_estimate_id", userId, estimatedDeliveryTime).Scan(&estimateId)

	if err != nil {
		return 0, err
	}

	return estimateId, nil
}
