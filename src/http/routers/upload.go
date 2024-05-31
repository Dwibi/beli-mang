package routers

import (
	v1uploadcontroller "github.com/Dwibi/beli-mang/src/http/controllers/upload"
	"github.com/Dwibi/beli-mang/src/http/middlewares"
)

func (r *Router) RegisterUpload() {
	v1uploadcontroller :=
		v1uploadcontroller.New(
			&v1uploadcontroller.V1Upload{
				DB:       r.DB,
				Uploader: r.Uploader,
			},
		)

	r.Router.Post("/image", middlewares.AuthMiddleware, v1uploadcontroller.Image)
	// r.Router.Get("/users/orders", middlewares.AuthMiddleware, v1ordercontroller.FindAll)
}
