package helpers

import (
	"errors"
	"net/url"
	"regexp"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

// Enum constants for entities.MerchantCategory
const (
	SmallRestaurant       entities.MerchantCategory = "SmallRestaurant"
	MediumRestaurant      entities.MerchantCategory = "MediumRestaurant"
	LargeRestaurant       entities.MerchantCategory = "LargeRestaurant"
	MerchandiseRestaurant entities.MerchantCategory = "MerchandiseRestaurant"
	BoothKiosk            entities.MerchantCategory = "BoothKiosk"
	ConvenienceStore      entities.MerchantCategory = "ConvenienceStore"
)

func ValidateMerchantCategory(category entities.MerchantCategory) error {
	switch category {
	case SmallRestaurant, MediumRestaurant, LargeRestaurant, MerchandiseRestaurant, BoothKiosk, ConvenienceStore:
		return nil
	default:
		return errors.New("invalid merchant category")
	}
}

const (
	Beverage   entities.ProductCategory = "Beverage"
	Food       entities.ProductCategory = "Food"
	Snack      entities.ProductCategory = "Snack"
	Condiments entities.ProductCategory = "Condiments"
	Additions  entities.ProductCategory = "Additions"
)

func ValidateItemsCategory(category entities.ProductCategory) error {
	switch category {
	case Beverage, Food, Snack, Condiments, Additions:
		return nil
	default:
		return errors.New("invalid product category")
	}
}

func ValidateURLWithDomain(u string) error {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return errors.New("invalid URL format")
	}

	// Regular expression to match a valid domain
	re := regexp.MustCompile(`\.[a-z]{2,}$`)
	if !re.MatchString(parsedURL.Host) {
		return errors.New("URL must contain a valid domain")
	}

	return nil
}

func ValidateLatAndLong(lat, long float64) bool {
	return lat >= -90 && lat <= 90 && long >= -180 && long <= 180
}

func ValidateStartingPoint(orders []entities.Orders) error {
	startingPoint := false

	for _, o := range orders {
		if startingPoint && o.IsStartingPoint {
			return errors.New("there's should be one isStartingPoint = true in orders")
		}

		if !startingPoint && o.IsStartingPoint {
			startingPoint = true
		}
	}

	if !startingPoint {
		return errors.New("there's should be one isStartingPoint = true in orders")
	}

	return nil
}
