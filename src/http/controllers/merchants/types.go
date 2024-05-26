package v1merchantscontroller

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type V1Merchant struct {
	DB *sql.DB
}

type IV1Merchant interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

func New(v1Merchant *V1Merchant) IV1Merchant {
	return v1Merchant
}
