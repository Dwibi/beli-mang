package entities

import "time"

type Merchants struct {
	Id        int
	Name      string
	ImageUrl  string
	Lat       string
	Long      string
	CreatedAt time.Time
}

type Location struct {
	Lat  float64 `json:"lat" validate:"required,number"`
	Long float64 `json:"long" validate:"required,number"`
}

type CreateMerchantParams struct {
	Name             string   `json:"name" validate:"required,min=2,max=30"`
	MerchantCategory string   `json:"merchantCategory" validate:"required"`
	ImageUrl         string   `json:"imageUrl" validate:"required,url"`
	Location         Location `json:"location" validate:"required"`
}
