package v1ordercontroller

import (
	"fmt"

	"github.com/Dwibi/beli-mang/src/helpers"
	orderrepository "github.com/Dwibi/beli-mang/src/repositories/order"
	orderitemrepository "github.com/Dwibi/beli-mang/src/repositories/order_items"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	ordersusecase "github.com/Dwibi/beli-mang/src/usecases/orders"
	"github.com/gofiber/fiber/v2"
)

type CreateOrderParams struct {
	CalculatedEstimateId string `json:"calculatedEstimateId" validate:"required"`
}

func (i V1Orders) Create(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)
	// Body parse
	var orderBody CreateOrderParams
	if err := c.BodyParser(&orderBody); err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate body
	if err := helpers.Validator.Struct(orderBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// usecase
	uu := ordersusecase.New(
		userrepository.New(i.DB),
		orderrepository.New(i.DB),
		orderitemrepository.New(i.DB),
	)

	orderIdStr, status, err := uu.Create(ordersusecase.CreateOrderParams{
		ReqUserId:            userId,
		CalculatedEstimateId: orderBody.CalculatedEstimateId,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	// result
	return c.Status(status).JSON(fiber.Map{
		"orderId": orderIdStr,
	})
}
