package orderrepository

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Dwibi/beli-mang/src/entities"
)

func (i sOrderRepository) FindMany(userId int, filters *entities.SearchOrderParams) (*[]entities.ResultListOrderItems, error) {
	query := `SELECT
		o.order_id AS order_id,
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

	orderMap := make(map[string]entities.ResultListOrderItems)
	for rows.Next() {
		var (
			orderId              string
			merchantId           string
			merchantName         string
			merchantImageUrl     string
			merchantCategory     string
			merchantLocationLat  float64
			merchantLocationLong float64
			merchantCreatedAt    time.Time
			itemId               string
			itemName             string
			productCategory      string
			itemPrice            int
			itemImageUrl         string
			quantity             int
			itemCreatedAt        time.Time
		)

		err := rows.Scan(
			&orderId,
			&merchantId,
			&merchantName,
			&merchantCategory,
			&merchantImageUrl,
			&merchantLocationLat,
			&merchantLocationLong,
			&merchantCreatedAt,
			&itemId,
			&itemName,
			&productCategory,
			&itemPrice,
			&quantity,
			&itemImageUrl,
			&itemCreatedAt,
		)
		if err != nil {
			return nil, err
		}

		item := entities.OrderItems{
			Id:              itemId,
			Name:            itemName,
			Price:           itemPrice,
			Quantity:        quantity,
			ImageUrl:        itemImageUrl,
			ProductCategory: productCategory,
			CreatedAt:       itemCreatedAt,
		}

		if order, exists := orderMap[orderId]; exists {
			order.Orders[0].Items = append(order.Orders[0].Items, item)
			orderMap[orderId] = order
		} else {
			merchant := entities.Merchants{
				Id:               merchantId,
				Name:             merchantName,
				ImageUrl:         merchantImageUrl,
				MerchantCategory: merchantCategory,
				CreatedAt:        merchantCreatedAt,
				Location: entities.Location{
					Lat:  merchantLocationLat,
					Long: merchantLocationLong,
				},
			}
			orderMap[orderId] = entities.ResultListOrderItems{
				OrderId: orderId,
				Orders: []entities.Order{
					{
						Merchant: merchant,
						Items:    []entities.OrderItems{item},
					},
				},
			}
		}
	}

	var orders []entities.ResultListOrderItems
	for _, order := range orderMap {
		orders = append(orders, order)
	}

	return &orders, nil
}
