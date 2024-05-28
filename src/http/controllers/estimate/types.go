package v1estimatescontroller

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type V1Estimates struct {
	DB *sql.DB
}

type IV1Estimates interface {
	Create(c *fiber.Ctx) error
	// FindAll(c *fiber.Ctx) error
}

func New(v1Estimates *V1Estimates) IV1Estimates {
	return v1Estimates
}
