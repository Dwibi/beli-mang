package v1merchantscontroller

import (
	"strconv"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	merchantrepository "github.com/Dwibi/beli-mang/src/repositories/merchants"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	merchantusecase "github.com/Dwibi/beli-mang/src/usecases/merchants"
	"github.com/gofiber/fiber/v2"
)

func (i V1Merchant) Create(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// Parse merchant body
	merchant := new(entities.CreateMerchantParams)
	if err := c.BodyParser(merchant); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Validate merchant payload
	if err := helpers.Validator.Struct(merchant); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := helpers.ValidateMerchantCategory(merchant.MerchantCategory); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := helpers.ValidateURLWithDomain(merchant.ImageUrl); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Create merchant
	uu := merchantusecase.New(
		userrepository.New(i.DB),
		merchantrepository.New(i.DB),
	)

	merchantId, status, err := uu.Create(&merchantusecase.CreateMerchantParams{
		UserId:   userId,
		Merchant: *merchant,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(fiber.Map{
		"merchantId": strconv.Itoa(merchantId),
	})
}
