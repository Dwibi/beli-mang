package merchantrepository

import "github.com/Dwibi/beli-mang/src/entities"

func (i sMerchantRepository) Create(m *entities.CreateMerchantParams) (int, error) {
	var merchantId int
	query := `INSERT INTO merchants (name, merchant_category, image_url, location_lat, location_long) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := i.DB.QueryRow(query, m.Name, m.MerchantCategory, m.ImageUrl, m.Location.Lat, m.Location.Long).Scan(&merchantId)

	if err != nil {
		return 0, err
	}

	return merchantId, nil
}
