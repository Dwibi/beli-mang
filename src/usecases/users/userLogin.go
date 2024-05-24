package userusecase

import (
	"errors"
	"os"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	"github.com/gofiber/fiber/v2"
)

func (i sUserUseCase) UserLogin(p entities.LoginParams) (string, int, error) {
	// Get user id and password
	user, err := i.userRepository.FindOne(&userrepository.FindOneParams{
		Username: p.Username,
		Role:     "user",
	})

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	if user == nil {
		return "", fiber.StatusNotFound, errors.New("user not found")
	}

	// Compare password
	isPasswordValid := helpers.ComparePassword(user.Password, p.Password)
	if !isPasswordValid {
		return "", fiber.StatusBadRequest, errors.New("password wrong")
	}

	// Create token
	token, err := helpers.CreateUserToken(&helpers.ParamCreateUser{
		SecretKey:       []byte(os.Getenv("JWT_SECRET")),
		ExpiredInMinute: 120,
		UserId:          user.Id,
	})

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	return token, fiber.StatusCreated, nil
}
