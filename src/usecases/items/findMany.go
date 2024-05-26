package itemsusecase

import (
	"errors"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/gofiber/fiber/v2"
)

type FindManyParams struct {
	UserId       int
	SearchParams entities.SearchItemsParams
}

func (i sItemsUseCase) FindManyParams(m *FindManyParams) (*entities.ItemsResult, int, error) {
	// Validate user who make a request is admin or not
	userRole, err := i.userRepository.FindUserRole(m.UserId)

	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	if userRole != "admin" {
		return nil, fiber.StatusUnauthorized, errors.New("only admin can this route")
	}

	// Find data merchant by repository
	merchants, err := i.itemsRepository.FindMany(&m.SearchParams)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return merchants, fiber.StatusOK, nil
}
