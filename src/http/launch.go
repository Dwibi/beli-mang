package http

import (
	"database/sql"

	"github.com/Dwibi/beli-mang/src/http/routers"
	"github.com/gofiber/fiber/v2"
)

type Http struct {
	DB *sql.DB
	// Uploader *manager.Uploader
}

type iHttp interface {
	Launch()
}

func New(Http *Http) iHttp {
	return Http
}

func (h *Http) Launch() {
	app := fiber.New()

	// Router
	api := app.Group("/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router := routers.New(&routers.Router{
		Router: api,
		DB:     h.DB,
	})

	router.RegisterUser()

	app.Listen(":8080")
}