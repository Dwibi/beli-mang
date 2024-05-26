package v1itemscontroller

import (
	"errors"
	"strconv"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	itemsrepository "github.com/Dwibi/beli-mang/src/repositories/items"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	itemsusecase "github.com/Dwibi/beli-mang/src/usecases/items"
	"github.com/gofiber/fiber/v2"
)

func (i V1Items) Create(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// Parse merchant body
	items := new(entities.CreateItemsParams)
	if err := c.BodyParser(items); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Validate merchant payload
	if err := helpers.Validator.Struct(items); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if items.Price < 0 {
		return fiber.NewError(fiber.StatusBadRequest, errors.New("price can't below 0").Error())
	}

	if err := helpers.ValidateItemsCategory(items.ProductCategory); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := helpers.ValidateURLWithDomain(items.ImageUrl); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Create items
	uu := itemsusecase.New(
		userrepository.New(i.DB),
		itemsrepository.New(i.DB),
	)

	itemId, status, err := uu.Create(&itemsusecase.CreateMerchantParams{
		UserId: userId,
		Items:  *items,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(fiber.Map{
		"itemId": strconv.Itoa(itemId),
	})
}
