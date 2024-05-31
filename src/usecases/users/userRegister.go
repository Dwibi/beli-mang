package userusecase

import (
	"errors"
	"os"

	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	"github.com/gofiber/fiber/v2"
)

func (i sUserUseCase) UserRegister(p entities.RegisterParams) (string, int, error) {
	// Validasi username udah di pake atau belum
	isUsernameExist, err := i.userRepository.IsExist(&userrepository.IsExistParams{
		Username: p.Username,
	})

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	if isUsernameExist {
		return "", fiber.StatusConflict, errors.New("username conflict with all types of users")
	}

	// Validasi email udah di pake atau belum
	isEmailExist, err := i.userRepository.IsExist(&userrepository.IsExistParams{
		Email: p.Email,
		Role:  "user",
	})

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	if isEmailExist {
		return "", fiber.StatusConflict, errors.New("email conflict with another user")
	}

	// Hashpassword
	hashedPassword, err := helpers.HashPassword(p.Password)
	if err != nil {
		return "", fiber.StatusInternalServerError, errors.New("failed to hash password")
	}

	// Create user
	userId, err := i.userRepository.Create(&userrepository.ParamsCreateUser{
		Username: p.Username,
		Email:    p.Email,
		Password: hashedPassword,
		Role:     "user",
	})
	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	// Create token

	token, err := helpers.CreateUserToken(&helpers.ParamCreateUser{
		SecretKey:       []byte(os.Getenv("JWT_SECRET")),
		ExpiredInMinute: 120,
		UserId:          userId,
	})

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	return token, fiber.StatusCreated, nil
}
