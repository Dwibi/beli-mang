package routers

import (
	v1merchantscontroller "github.com/Dwibi/beli-mang/src/http/controllers/merchants"
	"github.com/Dwibi/beli-mang/src/http/middlewares"
)

func (r *Router) RegisterMerchant() {
	v1merchantscontroller := v1merchantscontroller.New(
		&v1merchantscontroller.V1Merchant{
			DB: r.DB,
		},
	)

	r.Router.Post("/admin/merchants", middlewares.AuthMiddleware, v1merchantscontroller.Create)
	r.Router.Get("/admin/merchants", middlewares.AuthMiddleware, v1merchantscontroller.FindAll)
	r.Router.Get("/merchants/nearby/:coordinates", middlewares.AuthMiddleware, v1merchantscontroller.FindNearby)

}
