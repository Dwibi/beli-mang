package v1estimatescontroller

import (
	"errors"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	estimateitemsrepository "github.com/Dwibi/beli-mang/src/repositories/estimate_items"
	estimatesrepository "github.com/Dwibi/beli-mang/src/repositories/estimates"
	itemsrepository "github.com/Dwibi/beli-mang/src/repositories/items"
	merchantrepository "github.com/Dwibi/beli-mang/src/repositories/merchants"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	estimatesusecase "github.com/Dwibi/beli-mang/src/usecases/estimates"
	"github.com/gofiber/fiber/v2"
)

func (i V1Estimates) Create(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// Body Parse
	estimate := new(entities.CreateEstimateParams)

	if err := c.BodyParser(estimate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Validate
	if err := helpers.Validator.Struct(estimate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if isValidLatAndLong := helpers.ValidateLatAndLong(estimate.UserLocation.Lat, estimate.UserLocation.Long); !isValidLatAndLong {
		return fiber.NewError(fiber.StatusBadRequest, errors.New("invalid lat/long").Error())
	}

	if err := helpers.ValidateStartingPoint(estimate.Orders); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	uu := estimatesusecase.New(
		userrepository.New(i.DB),
		merchantrepository.New(i.DB),
		itemsrepository.New(i.DB),
		estimatesrepository.New(i.DB),
		estimateitemsrepository.New(i.DB),
	)

	result, status, err := uu.Create(estimatesusecase.CreateEstimateParams{
		UserId:        userId,
		EstimatesBody: *estimate,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(result)

}
