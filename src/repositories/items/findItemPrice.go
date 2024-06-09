package itemsrepository

import (
	"database/sql"
)

func (i sItemsRepository) FindItemPrice(id int) (int, error) {
	query := `
		SELECT 
			price
		FROM items 
		WHERE id = $1
	`

	var price int
	err := i.DB.QueryRow(query, id).Scan(
		&price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return price, nil
}
