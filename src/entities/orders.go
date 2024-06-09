package entities

import "time"

type SearchOrderParams struct {
	MerchantId       string
	Limit            int
	Offset           int
	Name             string
	MerchantCategory string
}

type OrderItems struct {
	Id              string    `json:"itemId"`
	Name            string    `json:"name"`
	ImageUrl        string    `json:"imageUrl"`
	ProductCategory string    `json:"productCategory"`
	Price           int       `json:"price"`
	Quantity        int       `json:"quantity"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Order struct {
	Merchant Merchants    `json:"merchant"`
	Items    []OrderItems `json:"items"`
}

type ResultListOrderItems struct {
	OrderId string  `json:"orderId"`
	Orders  []Order `json:"orders"`
}
