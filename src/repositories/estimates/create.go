package estimatesrepository

import "github.com/Dwibi/beli-mang/src/entities"

func (i sEstimatesRepository) Create(userId int, totalPrice, estimatedDeliveryTime float64) (*entities.ResultEstimate, error) {
	// var estimateId int
	var estimate entities.ResultEstimate
	err := i.DB.QueryRow("INSERT INTO estimates (user_id, total_price, estimated_delivery_time) VALUES ($1, $2, $3) RETURNING calculated_estimate_id, total_price, estimated_delivery_time", userId, totalPrice, estimatedDeliveryTime).Scan(&estimate.CalculatedEstimateId, &estimate.TotalPrice, &estimate.EstimatedDeliveryTimeInMinutes)

	if err != nil {
		return nil, err
	}

	return &estimate, nil
}
