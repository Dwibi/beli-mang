package entities

import "time"

type Items struct {
	Id              string    `json:"itemId"`
	Name            string    `json:"name"`
	ImageUrl        string    `json:"imageUrl"`
	ProductCategory string    `json:"productCategory"`
	Price           int       `json:"price"`
	CreatedAt       time.Time `json:"createdAt"`
}

type ProductCategory string

type CreateItemsParams struct {
	Name            string          `json:"name" validate:"required,min=2,max=30"`
	ProductCategory ProductCategory `json:"productCategory" validate:"required"`
	ImageUrl        string          `json:"imageUrl" validate:"required,url"`
	Price           int             `json:"price" validate:"required,number,min=1"`
}

type SearchItemsParams struct {
	ItemId          string
	Limit           int
	Offset          int
	Name            string
	ProductCategory string
	CreatedAt       string
}

type MetaTypes struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ItemsResult struct {
	Data []*Items  `json:"data"`
	Meta MetaTypes `json:"meta"`
}
