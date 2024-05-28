package routers

import (
	v1estimatescontroller "github.com/Dwibi/beli-mang/src/http/controllers/estimate"
	"github.com/Dwibi/beli-mang/src/http/middlewares"
)

func (r *Router) RegisterEstimates() {
	v1estimatescontroller := v1estimatescontroller.New(
		&v1estimatescontroller.V1Estimates{
			DB: r.DB,
		},
	)

	r.Router.Post("/users/estimate", middlewares.AuthMiddleware, v1estimatescontroller.Create)
}
