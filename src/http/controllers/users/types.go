package v1userscontroller

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type V1Users struct {
	DB *sql.DB
}

type IV1Users interface {
	AdminRegister(*fiber.Ctx) error
	AdminLogin(c *fiber.Ctx) error
	UserRegister(*fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
}

func New(v1Users *V1Users) IV1Users {
	return v1Users
}
