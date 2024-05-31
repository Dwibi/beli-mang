package v1uploadcontroller

import (
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gofiber/fiber/v2"
)

type V1Upload struct {
	DB       *sql.DB
	Uploader *manager.Uploader
}

type IV1Upload interface {
	// Create(c *fiber.Ctx) error
	// FindAll(c *fiber.Ctx) error
	Image(c *fiber.Ctx) error
}

func New(v1Upload *V1Upload) IV1Upload {
	return v1Upload
}
