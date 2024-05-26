package v1itemscontroller

import (
	"database/sql"
)

type V1Items struct {
	DB *sql.DB
}

type IV1Items interface {
	// Create(c *fiber.Ctx) error
	// FindAll(c *fiber.Ctx) error
}

func New(v1Items *V1Items) IV1Items {
	return v1Items
}
