package routers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Router fiber.Router
	DB     *sql.DB
	// Uploader *manager.Uploader
}

type iRoutes interface {
	// RegisterHello()
	RegisterUser()
	RegisterMerchant()
	// RegisterPatient()
	// RegisterRecord()
	// RegisterUpload()
}

func New(routes *Router) iRoutes {
	return routes
}
