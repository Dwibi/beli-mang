package merchantrepository

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sMerchantRepository) FindMany(filters *entities.SearchMerchantParams) (*entities.MerchantResult, error) {
	query := "SELECT id, name, merchant_category, image_url, location_lat, location_long, created_at FROM merchants WHERE 1=1"
	params := []interface{}{}

	n := (&entities.SearchMerchantParams{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.MerchantId != "" {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantId)
		}

		if filters.MerchantCategory != "" {
			conditions = append(conditions, "merchant_category = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantCategory)
		}

		if filters.Name != "" {
			conditions = append(conditions, "lower(name) LIKE lower($"+strconv.Itoa(len(params)+1)+")")
			params = append(params, "%"+filters.Name+"%")
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	if filters.CreatedAt != "" {
		if filters.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
		if filters.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		}
	} else {
		query += " ORDER BY created_at DESC"
	}

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

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding cat: %s", err)
		return nil, err
	}
	defer rows.Close()

	merchants := make([]*entities.Merchants, 0)
	meta := entities.MetaType{
		Limit:  filters.Limit,
		Offset: filters.Offset,
		Total:  len(merchants),
	}

	for rows.Next() {
		c := new(entities.Merchants)
		var MerchantIdStr string
		err := rows.Scan(&MerchantIdStr, &c.Name, &c.MerchantCategory, &c.ImageUrl, &c.Location.Lat, &c.Location.Long, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		c.Id = func() int { n, _ := strconv.Atoi(MerchantIdStr); return n }()
		merchants = append(merchants, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	result := entities.MerchantResult{
		Data: merchants,
		Meta: meta,
	}

	return &result, nil
}
