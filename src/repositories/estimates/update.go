package estimatesrepository

import "github.com/Dwibi/beli-mang/src/entities"

func (i sEstimatesRepository) Update(totalPrice float64) (*entities.ResultEstimate, error) {
	var result entities.ResultEstimate
	err := i.DB.QueryRow("UPDATE estimates SET total_price = $1 RETURNING calculated_estimate_id, total_price, estimated_delivery_time;", totalPrice).Scan(&result.CalculatedEstimateId, &result.TotalPrice, &result.EstimatedDeliveryTimeInMinutes)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
