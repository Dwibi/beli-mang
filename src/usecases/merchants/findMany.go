package merchantusecase

import (
	"errors"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/gofiber/fiber/v2"
)

type FindManyParams struct {
	UserId       int
	SearchParams entities.SearchMerchantParams
}

func (i sMerchantUseCase) FindManyParams(m *FindManyParams) (*entities.MerchantResult, int, error) {
	// Validate user who make a request is admin or not
	userRole, err := i.userRepository.FindUserRole(m.UserId)

	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	if userRole != "admin" {
		return nil, fiber.StatusUnauthorized, errors.New("only admin can this route")
	}

	// fmt.Println("limit usecases", m.SearchParams.Limit)
	// fmt.Println("offset usecases", m.SearchParams.Offset)

	// Find data merchant by repository
	merchants, err := i.merchantRepository.FindMany(&m.SearchParams)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return merchants, fiber.StatusOK, nil
}
