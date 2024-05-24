package v1userscontroller

import (
	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/Dwibi/beli-mang/src/helpers"
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	userusecase "github.com/Dwibi/beli-mang/src/usecases/users"
	"github.com/gofiber/fiber/v2"
)

func (i V1Users) AdminRegister(c *fiber.Ctx) error {
	user := new(entities.RegisterParams)
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := helpers.Validator.Struct(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	uu := userusecase.New(
		userrepository.New(i.DB),
	)

	token, status, err := uu.AdminRegister(entities.RegisterParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return fiber.NewError(status, err.Error())
	}

	return c.Status(status).JSON(fiber.Map{
		"token": token,
	})
}
