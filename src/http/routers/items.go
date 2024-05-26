package routers

import (
	v1merchantscontroller "github.com/Dwibi/beli-mang/src/http/controllers/merchants"
	"github.com/Dwibi/beli-mang/src/http/middlewares"
)

func (r *Router) RegisterItems() {
	v1merchantscontroller := v1merchantscontroller.New(
		&v1merchantscontroller.V1Merchant{
			DB: r.DB,
		},
	)

	r.Router.Post("/admin/merchants/:merchantId/items", middlewares.AuthMiddleware, v1merchantscontroller.Create)
	r.Router.Get("/admin/merchants/:merchantId/items", middlewares.AuthMiddleware, v1merchantscontroller.FindAll)
	// r.Router.Post("/admin/login", v1userscontroller.AdminLogin)
	// r.Router.Post("/user/register", v1userscontroller.UserRegister)
	// r.Router.Post("/user/login", v1userscontroller.UserLogin)

}
