package merchantrepository

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sMerchantRepository) FindDistance(latitude, longitude float64, ids []int) (*entities.FindDistanceResult, error) {
	query := `
    SELECT id, location_lat, location_long ,(
        6371 * 2 * asin(sqrt(
            power(sin(radians(m.location_lat - $1) / 2), 2) +
            cos(radians($1)) * cos(radians(m.location_lat)) * 
            power(sin(radians(m.location_long - $2) / 2), 2)
        ))
    ) AS distance
    FROM merchants m
    WHERE id IN`

	params := []interface{}{latitude, longitude}

	// Declaring conditions slice outside the conditional block
	conditions := []string{}

	for _, value := range ids {
		conditions = append(conditions, "$"+strconv.Itoa(len(params)+1))
		params = append(params, value)
	}

	query += "("
	query += strings.Join(conditions, ", ")
	query += ") ORDER BY distance LIMIT 1;"

	// fmt.Println("query : ", query)

	var result entities.FindDistanceResult
	err := i.DB.QueryRow(query, params...).Scan(&result.Id, &result.Lat, &result.Long, &result.Distance)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
