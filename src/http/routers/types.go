package routers

import (
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Router   fiber.Router
	DB       *sql.DB
	Uploader *manager.Uploader
}

type iRoutes interface {
	RegisterUser()
	RegisterMerchant()
	RegisterItems()
	RegisterEstimates()
	RegisterOrders()
	RegisterUpload()
}

func New(routes *Router) iRoutes {
	return routes
}
