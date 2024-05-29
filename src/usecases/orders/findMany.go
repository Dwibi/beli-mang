package ordersusecase

import (
	"errors"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/gofiber/fiber/v2"
)

type FindManyParams struct {
	UserId       int
	SearchParams entities.SearchOrderParams
}

func (i sOrdersUseCase) FindMany(m *FindManyParams) (*entities.ResultListOrderItems, int, error) {
	// Validate user who make a request is admin or not
	userRole, err := i.userRepository.FindUserRole(m.UserId)

	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	if userRole != "user" {
		return nil, fiber.StatusUnauthorized, errors.New("only user can this route")
	}

	// Find data merchant by repository
	result, err := i.orderRepository.FindMany(m.UserId, &m.SearchParams)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return result, fiber.StatusOK, nil
}
