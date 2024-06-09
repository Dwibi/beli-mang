package merchantrepository

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sMerchantRepository) FindOne(id int) (*entities.Merchants, error) {
	query := `
		SELECT 
			id, 
			name, 
			merchant_category, 
			image_url, 
			location_lat, 
			location_long, 
			created_at 
		FROM merchants 
		WHERE id = $1
	`

	var merchant entities.Merchants
	err := i.DB.QueryRow(query, id).Scan(
		&merchant.Id,
		&merchant.Name,
		&merchant.MerchantCategory,
		&merchant.ImageUrl,
		&merchant.Location.Lat,
		&merchant.Location.Long,
		&merchant.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &merchant, nil
}
