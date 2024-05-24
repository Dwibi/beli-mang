package v1merchantscontroller

import (
	"github.com/Dwibi/beli-mang/src/entities"
	"github.com/gofiber/fiber/v2"
)

func (i V1Merchant) Create(c *fiber.Ctx) error {
	// Parse merchant body
	merchant := new(entities.CreateMerchantParams)
	if err := c.BodyParser(merchant); err != nil {
		// return fiber.NewError()
	}
	// Validate merchant payload
	// Create merchant
	// Return merchant id
}
