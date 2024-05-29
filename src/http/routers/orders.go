package routers

import (
	v1ordercontroller "github.com/Dwibi/beli-mang/src/http/controllers/orders"
	"github.com/Dwibi/beli-mang/src/http/middlewares"
)

func (r *Router) RegisterOrders() {
	v1ordercontroller := v1ordercontroller.New(
		&v1ordercontroller.V1Orders{
			DB: r.DB,
		},
	)

	r.Router.Post("/users/orders", middlewares.AuthMiddleware, v1ordercontroller.Create)
	r.Router.Get("/users/orders", middlewares.AuthMiddleware, v1ordercontroller.FindAll)
}
