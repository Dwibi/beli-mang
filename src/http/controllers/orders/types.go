package v1ordercontroller

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type V1Orders struct {
	DB *sql.DB
}

type IV1Orders interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

func New(v1Orders *V1Orders) IV1Orders {
	return v1Orders
}
