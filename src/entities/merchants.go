package entities

import (
	"time"
)

type Merchants struct {
	Id               string    `json:"merchantId"`
	Name             string    `json:"name"`
	ImageUrl         string    `json:"imageUrl"`
	MerchantCategory string    `json:"merchantCategory"`
	Location         Location  `json:"location"`
	CreatedAt        time.Time `json:"createdAt"`
}

type Location struct {
	Lat  float64 `json:"lat" validate:"required,number"`
	Long float64 `json:"long" validate:"required,number"`
}

type MerchantCategory string

type CreateMerchantParams struct {
	Name             string           `json:"name" validate:"required,min=2,max=30"`
	MerchantCategory MerchantCategory `json:"merchantCategory" validate:"required"`
	ImageUrl         string           `json:"imageUrl" validate:"required,url"`
	Location         Location         `json:"location" validate:"required"`
}

type SearchMerchantParams struct {
	MerchantId       string
	Limit            int
	Offset           int
	Name             string
	MerchantCategory string
	CreatedAt        string
}

type SearchNearbyMerchantParams struct {
	MerchantId       string
	Limit            int
	Offset           int
	Name             string
	MerchantCategory string
}
type MetaType struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
type MerchantResult struct {
	Data []*Merchants `json:"data"`
	Meta MetaType     `json:"meta"`
}

type FindDistanceResult struct {
	Id       int
	Lat      float64
	Long     float64
	Distance float64
}

type GetNearbyMerchantQueries struct {
	MerchantId       string  `db:"id" json:"merchantId" query:"merchantId"`
	Limit            int     `json:"limit" query:"limit"`
	Offset           int     `json:"offset" query:"offset"`
	Name             string  `db:"name" json:"name" query:"name"`
	MerchantCategory string  `db:"merchant_category" json:"merchantCategory" query:"merchantCategory"`
	Latitude         float64 `db:"latitude" json:"lat"`
	Longitude        float64 `db:"longitude" json:"long"`
}

type GetNearbyMerchantResponse struct {
	Merchant Merchants `json:"merchant"`
	Items    []Items   `json:"items"`
}
