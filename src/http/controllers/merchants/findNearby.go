package v1merchantscontroller

import (
	"strconv"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
	merchantrepository "github.com/Dwibi/beli-mang/src/repositories/merchants"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	merchantusecase "github.com/Dwibi/beli-mang/src/usecases/merchants"
	"github.com/gofiber/fiber/v2"
)

func (i V1Merchant) FindNearby(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// Acessign user location
	location := c.Params("coordinates")
	coords := strings.Split(location, ",")
	if len(coords) != 2 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid coordinates format. Expected format: lat,long")
	}

	latStr := coords[0]
	longStr := coords[1]

	// Convert latitude and longitude to float64
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid latitude value")
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid longitude value")
	}

	// get queries
	q := c.Queries()
	filters := new(entities.SearchNearbyMerchantParams)

	if q["merchantId"] != "" {
		filters.MerchantId = q["merchantId"]
	}

	if q["limit"] != "" {
		limit, err := strconv.Atoi(q["limit"])
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "limit must be number")
		}
		filters.Limit = limit
	}

	if q["offset"] != "" {
		offset, err := strconv.Atoi(q["offset"])
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "offset must be number")
		}
		filters.Limit = offset
	}

	if q["name"] != "" {
		filters.Name = q["name"]
	}

	if q["merchantCategory"] != "" {
		// TODO: validasi categorynya
		filters.MerchantCategory = q["merchantCategory"]
	}

	// Get merchant
	uu := merchantusecase.New(
		userrepository.New(i.DB),
		merchantrepository.New(i.DB),
	)

	result, status, err := uu.FindNearby(&merchantusecase.FindNearbyParams{
		UserId:       userId,
		Latitude:     lat,
		Longtitude:   long,
		SearchParams: *filters,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(result)
}
