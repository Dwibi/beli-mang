package http

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/http/middlewares"
	"github.com/Dwibi/beli-mang/src/http/routers"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gofiber/fiber/v2"
)

type Http struct {
	DB       *sql.DB
	Uploader *manager.Uploader
}

type iHttp interface {
	Launch()
}

func New(Http *Http) iHttp {
	return Http
}

func (h *Http) Launch() {
	app := fiber.New()

	app.Get("/", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router := routers.New(&routers.Router{
		Router:   app,
		DB:       h.DB,
		Uploader: h.Uploader,
	})

	router.RegisterUser()
	router.RegisterMerchant()
	router.RegisterItems()
	router.RegisterEstimates()
	router.RegisterOrders()
	router.RegisterUpload()

	app.Listen(":8080")
}
