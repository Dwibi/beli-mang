package v1ordercontroller

import (
	"strconv"

	"github.com/Dwibi/beli-mang/src/entities"
	orderrepository "github.com/Dwibi/beli-mang/src/repositories/order"
	orderitemrepository "github.com/Dwibi/beli-mang/src/repositories/order_items"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	ordersusecase "github.com/Dwibi/beli-mang/src/usecases/orders"
	"github.com/gofiber/fiber/v2"
)

func (i V1Orders) FindAll(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// get queries
	q := c.Queries()
	filters := new(entities.SearchOrderParams)

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
	uu := ordersusecase.New(
		userrepository.New(i.DB),
		orderrepository.New(i.DB),
		orderitemrepository.New(i.DB),
	)

	result, status, err := uu.FindMany(&ordersusecase.FindManyParams{
		UserId:       userId,
		SearchParams: *filters,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(result)
}
