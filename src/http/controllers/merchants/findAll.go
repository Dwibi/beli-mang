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

func (i V1Merchant) FindAll(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// get queries
	q := c.Queries()
	filters := new(entities.SearchMerchantParams)

	// fmt.Println("limit controller", q["limit"])
	// fmt.Println("offset controller", q["offset"])

	if q["merchantId"] != "" {
		filters.MerchantId = q["merchantId"]
	}

	if q["limit"] != "" {
		limitInt, err := strconv.Atoi(q["limit"])
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "limit must be number")
		}
		filters.Limit = limitInt
	}

	if q["offset"] != "" {
		offsetInt, err := strconv.Atoi(q["offset"])
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "offset must be number")
		}
		filters.Offset = offsetInt
	}

	if q["name"] != "" {
		filters.Name = q["name"]
	}

	if q["merchantCategory"] != "" {
		// TODO: validasi categorynya
		filters.MerchantCategory = q["merchantCategory"]
	}

	if q["createdAt"] != "" {
		lowCreatedAt := strings.ToLower(q["createdAt"])
		if lowCreatedAt == "asc" || lowCreatedAt == "desc" {
			filters.CreatedAt = lowCreatedAt
		}
	}

	// Get merchant
	uu := merchantusecase.New(
		userrepository.New(i.DB),
		merchantrepository.New(i.DB),
	)

	result, status, err := uu.FindManyParams(&merchantusecase.FindManyParams{
		UserId:       userId,
		SearchParams: *filters,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(result)
}
