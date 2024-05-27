package routers

import (
	v1itemscontroller "github.com/Dwibi/beli-mang/src/http/controllers/items"
	"github.com/Dwibi/beli-mang/src/http/middlewares"
)

func (r *Router) RegisterItems() {
	v1itemscontroller := v1itemscontroller.New(
		&v1itemscontroller.V1Items{
			DB: r.DB,
		},
	)

	r.Router.Post("/admin/merchants/:merchantId/items", middlewares.AuthMiddleware, v1itemscontroller.Create)
	r.Router.Get("/admin/merchants/:merchantId/items", middlewares.AuthMiddleware, v1itemscontroller.FindAll)
	// r.Router.Post("/admin/login", v1userscontroller.AdminLogin)
	// r.Router.Post("/user/register", v1userscontroller.UserRegister)
	// r.Router.Post("/user/login", v1userscontroller.UserLogin)

}
