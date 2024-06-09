package itemsusecase

import (
	"errors"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/gofiber/fiber/v2"
)

type CreateMerchantParams struct {
	UserId     int
	MerchantId int
	Items      entities.CreateItemsParams
}

func (i sItemsUseCase) Create(m *CreateMerchantParams) (int, int, error) {
	// Validate user who make a request is admin or not
	userRole, err := i.userRepository.FindUserRole(m.UserId)

	if err != nil {
		return 0, fiber.StatusInternalServerError, err
	}

	if userRole != "admin" {
		return 0, fiber.StatusUnauthorized, errors.New("only admin can this route")
	}

	// fmt.Println(m.Items)
	// Create data di repository
	merchantId, err := i.itemsRepository.Create(m.MerchantId, &m.Items)

	if err != nil {
		return 0, fiber.StatusInternalServerError, err
	}

	// Returning merchant id
	return merchantId, fiber.StatusCreated, nil
}
