package itemsrepository

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sItemsRepository) FindMany(merchantId int, filters *entities.SearchItemsParams) (*entities.ItemsResult, error) {
	query := "SELECT id, name, product_category, image_url, price, created_at FROM items WHERE 1=1"
	params := []interface{}{}

	// Declaring conditions slice outside the conditional block
	conditions := []string{}

	n := (&entities.SearchMerchantParams{})

	if !reflect.DeepEqual(filters, n) {
		if filters.ItemId != "" {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.ItemId)
		}

		if filters.ProductCategory != "" {
			conditions = append(conditions, "product_category = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.ProductCategory)
		}

		if filters.Name != "" {
			conditions = append(conditions, "lower(name) LIKE lower($"+strconv.Itoa(len(params)+1)+")")
			params = append(params, "%"+filters.Name+"%")
		}

		conditions = append(conditions, "merchant_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, merchantId)

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	// Count query to get the total number of items
	countQuery := "SELECT COUNT(*) FROM items WHERE 1=1"
	if len(conditions) > 0 {
		countQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Execute the count query
	var total int
	err := i.DB.QueryRow(countQuery, params...).Scan(&total)
	if err != nil {
		log.Printf("Error counting items: %s", err)
		return nil, err
	}

	// Append ORDER BY clause for the main query
	if filters.CreatedAt != "" {
		if filters.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		} else if filters.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		}
	} else {
		query += " ORDER BY created_at DESC"
	}

	// Append LIMIT and OFFSET clauses for the main query
	if filters.Limit == 0 {
		filters.Limit = 5
	}
	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	// Execute the main query
	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding items: %s", err)
		return nil, err
	}
	defer rows.Close()

	items := make([]*entities.Items, 0)
	meta := entities.MetaTypes{
		Limit:  filters.Limit,
		Offset: filters.Offset,
		Total:  total,
	}

	for rows.Next() {
		c := new(entities.Items)
		err := rows.Scan(&c.Id, &c.Name, &c.ProductCategory, &c.ImageUrl, &c.Price, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	result := entities.ItemsResult{
		Data: items,
		Meta: meta,
	}

	return &result, nil
}
