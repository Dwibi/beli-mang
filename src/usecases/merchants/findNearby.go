package merchantusecase

import (
	"errors"

	"github.com/Dwibi/beli-mang/src/entities"
	merchantrepository "github.com/Dwibi/beli-mang/src/repositories/merchants"
	"github.com/gofiber/fiber/v2"
)

type FindNearbyParams struct {
	UserId       int
	Latitude     float64
	Longtitude   float64
	SearchParams entities.SearchNearbyMerchantParams
}

func (i sMerchantUseCase) FindNearby(m *FindNearbyParams) (*merchantrepository.ResultFindNearby, int, error) {
	// Validate user who make a request is admin or not
	userRole, err := i.userRepository.FindUserRole(m.UserId)

	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	if userRole != "user" {
		return nil, fiber.StatusUnauthorized, errors.New("only user can this route")
	}

	// Find data merchant by repository
	result, err := i.merchantRepository.FindNearby(m.Latitude, m.Longtitude, &m.SearchParams)
	if err != nil {
		return nil, fiber.StatusInternalServerError, err
	}

	return result, fiber.StatusOK, nil
}
