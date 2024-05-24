package v1merchantscontroller

import (
	"database/sql"
)

type V1Merchant struct {
	DB *sql.DB
}

type IV1Merchant interface {
	// AdminRegister(*fiber.Ctx) error
	// AdminLogin(c *fiber.Ctx) error
	// UserRegister(*fiber.Ctx) error
	// UserLogin(c *fiber.Ctx) error
}

func New(v1Merchant *V1Merchant) IV1Merchant {
	return v1Merchant
}
