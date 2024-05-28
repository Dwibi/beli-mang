package entities

type UserLocation struct {
	Lat  float64 `json:"lat" validate:"required,number"`
	Long float64 `json:"long" validate:"required,number"`
}

type ItemsEstimate struct {
	ItemId   string `json:"itemId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,number"`
}

type Orders struct {
	MerchantId      string          `json:"merchantId"`
	IsStartingPoint bool            `json:"isStartingPoint" validate:"required,boolean"`
	Items           []ItemsEstimate `json:"items"`
}

type CreateEstimateParams struct {
	UserLocation UserLocation `json:"userLocation"`
	Orders       []Orders     `json:"orders"`
}

type ResultEstimate struct {
	TotalPrice                     float64 `json:"totalPrice"`
	EstimatedDeliveryTimeInMinutes float64 `json:"estimatedDeliveryTimeInMinutes"`
	CalculatedEstimateId           string  `json:"calculatedEstimateId"`
}
