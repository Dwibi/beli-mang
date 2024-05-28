package estimateitemsrepository

import (
	"fmt"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sEstimateItemsRepository) Create(calculatedEstimateId int, orders []entities.Orders) error {
	query := "INSERT INTO estimate_items (calculated_estimate_id, merchant_id, item_id, quantity, is_starting_point) VALUES "

	// Create a slice to hold the value strings and a slice to hold the arguments
	var valueStrings []string
	var valueArgs []interface{}
	argIndex := 1

	for _, order := range orders {
		for _, item := range order.Items {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", argIndex, argIndex+1, argIndex+2, argIndex+3, argIndex+4))
			valueArgs = append(valueArgs, calculatedEstimateId, order.MerchantId, item.ItemId, item.Quantity, order.IsStartingPoint)
			argIndex += 5
		}
	}

	// Join the value strings and append to the query
	query += strings.Join(valueStrings, ", ")

	// Execute the query
	_, err := i.DB.Exec(query, valueArgs...)
	return err
}
