package ordersusecase

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateOrderParams struct {
	ReqUserId            int
	CalculatedEstimateId string
}

func (i sOrdersUseCase) Create(o CreateOrderParams) (string, int, error) {
	userRole, err := i.userRepository.FindUserRole(o.ReqUserId)

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	if userRole != "user" {
		return "", fiber.StatusUnauthorized, errors.New("only user can this route")
	}

	// create orders
	orderId, err := i.orderRepository.Create(o.ReqUserId, o.CalculatedEstimateId)

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}
	// Create order items
	if err := i.orderItemRepository.Create(orderId, o.CalculatedEstimateId); err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	orderIdStr := strconv.Itoa(orderId)
	return orderIdStr, fiber.StatusCreated, nil

}
