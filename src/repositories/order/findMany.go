package orderrepository

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sOrderRepository) FindMany(userId int, filters *entities.SearchOrderParams) (*entities.ResultListOrderItems, error) {
	query := `SELECT
	oi.order_id AS order_id,
	m.id AS merchant_id,
    m.name AS merchant_name,
    m.merchant_category AS merchant_category,
    m.image_url AS merchant_image_url,
    m.location_lat AS merchant_location_lat,
    m.location_long AS merchant_location_long,
    m.created_at AS merchant_created_at,
    i.id AS item_id,
    i.name AS item_name,
    i.product_category AS item_product_category,
    i.price AS item_price,
	oi.quantity AS quantity,
    i.image_url AS item_image_url,
    i.created_at AS item_created_at
	FROM orders o
	JOIN order_items oi ON o.order_id = oi.order_id 
	JOIN merchants m ON oi.merchant_id = m.id 
	JOIN items i ON oi.item_id = i.id WHERE o.user_id = $1`
	params := []interface{}{userId}
	conditions := []string{}

	if filters != nil {
		if filters.MerchantId != "" {
			conditions = append(conditions, "m.id::text = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantId)
		}

		if filters.MerchantCategory != "" {
			conditions = append(conditions, "m.merchant_category = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantCategory)
		}

		if filters.Name != "" {
			conditions = append(conditions, "(m.name ILIKE $"+strconv.Itoa(len(params)+1)+" OR i.name ILIKE $"+strconv.Itoa(len(params)+2)+")")
			params = append(params, "%"+filters.Name+"%")
			params = append(params, "%"+filters.Name+"%")
		}

		if len(conditions) > 0 {
			query += " AND " + strings.Join(conditions, " AND ")
		}
	}

	query += " ORDER BY o.created_at DESC"

	if filters.Limit == 0 {
		filters.Limit = 5
	}
	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset != 0 {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}
	defer rows.Close()

	var result entities.ResultListOrderItems
	merchantOrderItemsMap := make(map[int]*entities.Order)

	for rows.Next() {
		var (
			orderID              int
			merchantID           int
			merchantName         string
			merchantCategory     string
			merchantImageUrl     string
			merchantLocationLat  float64
			merchantLocationLong float64
			merchantCreatedAt    time.Time
			itemID               int
			itemName             string
			itemProductCategory  string
			itemImageUrl         string
			itemPrice            int
			itemQuantity         int
			itemCreatedAt        time.Time
		)

		err := rows.Scan(&orderID, &merchantID, &merchantName, &merchantCategory, &merchantImageUrl, &merchantLocationLat,
			&merchantLocationLong, &merchantCreatedAt, &itemID, &itemName, &itemProductCategory,
			&itemPrice, &itemQuantity, &itemImageUrl, &itemCreatedAt)

		if err != nil {
			return nil, err
		}

		if _, exists := merchantOrderItemsMap[merchantID]; !exists {
			merchantOrderItemsMap[merchantID] = &entities.Order{
				OrderId: orderID,
				Merchant: entities.Merchants{
					Id:               strconv.Itoa(merchantID),
					Name:             merchantName,
					MerchantCategory: merchantCategory,
					ImageUrl:         merchantImageUrl,
					Location: entities.Location{
						Lat:  merchantLocationLat,
						Long: merchantLocationLong,
					},
					CreatedAt: merchantCreatedAt,
				},
				Items: []entities.OrderItems{},
			}
		}

		if filters != nil && strings.Contains(strings.ToLower(itemName), strings.ToLower(filters.Name)) {
			merchantOrderItemsMap[merchantID].Items = append(merchantOrderItemsMap[merchantID].Items, entities.OrderItems{
				Id:              strconv.Itoa(itemID),
				Name:            itemName,
				ProductCategory: itemProductCategory,
				ImageUrl:        itemImageUrl,
				Price:           itemPrice,
				Quantity:        itemQuantity,
				CreatedAt:       itemCreatedAt,
			})
		} else if filters == nil {
			merchantOrderItemsMap[merchantID].Items = append(merchantOrderItemsMap[merchantID].Items, entities.OrderItems{
				Id:              strconv.Itoa(itemID),
				Name:            itemName,
				ProductCategory: itemProductCategory,
				ImageUrl:        itemImageUrl,
				Price:           itemPrice,
				Quantity:        itemQuantity,
				CreatedAt:       itemCreatedAt,
			})
		}

	}

	for _, merchantOrderItem := range merchantOrderItemsMap {
		result.Data = append(result.Data, *merchantOrderItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &result, nil
}
