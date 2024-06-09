package merchantrepository

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Dwibi/beli-mang/src/entities"
)

type MerchantItem struct {
	Merchant entities.Merchants `json:"merchant"`
	Items    []entities.Items   `json:"items"`
}

type ResultFindNearby struct {
	Data []entities.GetNearbyMerchantResponse `json:"data"`
}

func (i sMerchantRepository) FindNearby(Latitude, Longitude float64, filters *entities.SearchNearbyMerchantParams) (*ResultFindNearby, error) {
	var merchants []entities.GetNearbyMerchantResponse

	var distance float64

	query := `
    SELECT id, name, merchant_category, image_url, location_lat, location_long, created_at,
        (acos(
            cos(radians($1)) * cos(radians(location_lat)) *
            cos(radians(location_long) - radians($2)) +
            sin(radians($1)) * sin(radians(location_lat))
        )) AS distance
    FROM merchants` //No need for earth constant because they all would be multiplied with the same constant

	params := []interface{}{Latitude, Longitude}
	conditions := []string{}

	if filters != nil {
		if filters.MerchantId != "" {
			conditions = append(conditions, "id::text = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantId)
		}

		if filters.MerchantCategory != "" {
			conditions = append(conditions, "merchant_category = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantCategory)
		}

		if filters.Name != "" {
			conditions = append(conditions, "name ILIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+filters.Name+"%")
		}

		if len(conditions) > 0 {
			query += " WHERE " + strings.Join(conditions, " AND ")
		}
	}

	query += " ORDER BY distance"

	if filters.Limit == 0 {
		filters.Limit = 5
	}
	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset != 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	// fmt.Println(query)

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		merchant := entities.GetNearbyMerchantResponse{}
		var createdAt time.Time
		err := rows.Scan(&merchant.Merchant.Id, &merchant.Merchant.Name, &merchant.Merchant.MerchantCategory, &merchant.Merchant.ImageUrl, &merchant.Merchant.Location.Lat, &merchant.Merchant.Location.Long, &createdAt, &distance)
		if err != nil {
			return nil, err
		}

		getItemsQuery := "SELECT id, name, product_category, price, image_url, created_at FROM items WHERE merchant_id = $1"

		item_rows, err := i.DB.Query(getItemsQuery, merchant.Merchant.Id)
		if err != nil {
			return nil, err
		}

		for item_rows.Next() {
			item := entities.Items{}

			err := item_rows.Scan(&item.Id, &item.Name, &item.ProductCategory, &item.Price, &item.ImageUrl, &item.CreatedAt)
			if err != nil {
				return nil, err
			}

			merchant.Items = append(merchant.Items, item)

		}

		merchants = append(merchants, merchant)
	}

	result := ResultFindNearby{
		Data: merchants,
	}

	return &result, nil

}

// func (i sMerchantRepository) FindNearby(Latitude, Longitude float64, filters *entities.SearchNearbyMerchantParams) (*ResultFindNearby, error) {
// 	query := `
// 	SELECT
//     m.id AS merchant_id,
//     m.name AS merchant_name,
//     m.merchant_category AS merchant_category,
//     m.image_url AS merchant_image_url,
//     m.location_lat AS merchant_location_lat,
//     m.location_long AS merchant_location_long,
//     m.created_at AS merchant_created_at,
//     i.id AS item_id,
//     i.name AS item_name,
//     i.product_category AS item_product_category,
//     i.image_url AS item_image_url,
//     i.price AS item_price,
//     i.created_at AS item_created_at,
//     (
//         6371 * 2 * asin(sqrt(
//             power(sin(radians(m.location_lat - $1) / 2), 2) +
//             cos(radians($1)) * cos(radians(m.location_lat)) *
//             power(sin(radians(m.location_long - $2) / 2), 2)
//         ))
//     ) AS distance
// FROM merchants m
// JOIN items i ON m.id = i.merchant_id
// WHERE 1=1`

// 	params := []interface{}{Latitude, Longitude}
// 	conditions := []string{}

// 	if filters != nil {
// 		if filters.MerchantId != "" {
// 			conditions = append(conditions, "m.id::text = $"+strconv.Itoa(len(params)+1))
// 			params = append(params, filters.MerchantId)
// 		}

// 		if filters.MerchantCategory != "" {
// 			conditions = append(conditions, "m.merchant_category = $"+strconv.Itoa(len(params)+1))
// 			params = append(params, filters.MerchantCategory)
// 		}

// 		if filters.Name != "" {
// 			conditions = append(conditions, "(m.name ILIKE $"+strconv.Itoa(len(params)+1)+" OR i.name ILIKE $"+strconv.Itoa(len(params)+2)+")")
// 			params = append(params, "%"+filters.Name+"%")
// 			params = append(params, "%"+filters.Name+"%")
// 		}

// 		if len(conditions) > 0 {
// 			query += " AND " + strings.Join(conditions, " AND ")
// 		}
// 	}

// 	query += " ORDER BY distance"

// 	if filters.Limit == 0 {
// 		filters.Limit = 5
// 	}
// 	query += " LIMIT $" + strconv.Itoa(len(params)+1)
// 	params = append(params, filters.Limit)

// 	if filters.Offset != 0 {
// 		query += " OFFSET $" + strconv.Itoa(len(params)+1)
// 		params = append(params, filters.Offset)
// 	}

// 	rows, err := i.DB.Query(query, params...)
// 	if err != nil {
// 		log.Printf("Error executing query: %s", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var result ResultFindNearby
// 	merchantItemsMap := make(map[int]*MerchantItem)

// 	for rows.Next() {
// 		var (
// 			merchantID           int
// 			merchantName         string
// 			merchantCategory     string
// 			merchantImageUrl     string
// 			merchantLocationLat  float64
// 			merchantLocationLong float64
// 			merchantCreatedAt    time.Time
// 			itemID               int
// 			itemName             string
// 			itemProductCategory  string
// 			itemImageUrl         string
// 			itemPrice            int
// 			itemCreatedAt        time.Time
// 			distance             float64
// 		)

// 		err := rows.Scan(&merchantID, &merchantName, &merchantCategory, &merchantImageUrl, &merchantLocationLat,
// 			&merchantLocationLong, &merchantCreatedAt, &itemID, &itemName, &itemProductCategory, &itemImageUrl,
// 			&itemPrice, &itemCreatedAt, &distance)

// 		if err != nil {
// 			return nil, err
// 		}

// 		// Debugging: Print each merchant's information
// 		fmt.Printf("Merchant ID: %d, Name: %s, Distance: %.2f km\n", merchantID, merchantName, distance)

// 		if _, exists := merchantItemsMap[merchantID]; !exists {
// 			merchantItemsMap[merchantID] = &MerchantItem{
// 				Merchant: entities.Merchants{
// 					Id:               strconv.Itoa(merchantID),
// 					Name:             merchantName,
// 					MerchantCategory: merchantCategory,
// 					ImageUrl:         merchantImageUrl,
// 					Location: entities.Location{
// 						Lat:  merchantLocationLat,
// 						Long: merchantLocationLong,
// 					},
// 					CreatedAt: merchantCreatedAt,
// 				},
// 				Items: []entities.Items{},
// 			}
// 		}

// 		// Add items to the merchant
// 		merchantItemsMap[merchantID].Items = append(merchantItemsMap[merchantID].Items, entities.Items{
// 			Id:              strconv.Itoa(itemID),
// 			Name:            itemName,
// 			ProductCategory: itemProductCategory,
// 			ImageUrl:        itemImageUrl,
// 			Price:           itemPrice,
// 			CreatedAt:       itemCreatedAt,
// 		})
// 	}

// 	for _, merchantItem := range merchantItemsMap {
// 		result.Data = append(result.Data, *merchantItem)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
