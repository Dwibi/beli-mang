package v1itemscontroller

import (
	"strconv"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
	itemsrepository "github.com/Dwibi/beli-mang/src/repositories/items"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	itemsusecase "github.com/Dwibi/beli-mang/src/usecases/items"
	"github.com/gofiber/fiber/v2"
)

func (i V1Items) FindAll(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	merchantId, _ := strconv.Atoi(c.Params("merchantId"))

	// get queries
	q := c.Queries()
	filters := new(entities.SearchItemsParams)

	if q["itemId"] != "" {
		filters.ItemId = q["itemId"]
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

	if q["productCategory"] != "" {
		// TODO: validasi categorynya
		filters.ProductCategory = q["productCategory"]
	}

	if q["createdAt"] != "" {
		lowCreatedAt := strings.ToLower(q["createdAt"])
		if lowCreatedAt == "asc" || lowCreatedAt == "desc" {
			filters.CreatedAt = lowCreatedAt
		}
	}

	// Get merchant
	uu := itemsusecase.New(
		userrepository.New(i.DB),
		itemsrepository.New(i.DB),
	)

	result, status, err := uu.FindManyParams(&itemsusecase.FindManyParams{
		UserId:       userId,
		MerchantId:   merchantId,
		SearchParams: *filters,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(result)
}
