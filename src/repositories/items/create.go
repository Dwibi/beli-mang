package itemsrepository

import "github.com/Dwibi/beli-mang/src/entities"

func (i sItemsRepository) Create(merchantId int, m *entities.CreateItemsParams) (int, error) {
	var itemsId int
	query := `INSERT INTO items (name, product_category, price, image_url, merchant_id) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := i.DB.QueryRow(query, m.Name, m.ProductCategory, m.Price, m.ImageUrl, merchantId).Scan(&itemsId)

	if err != nil {
		return 0, err
	}

	return itemsId, nil
}
