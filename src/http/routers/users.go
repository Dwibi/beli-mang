package routers

import v1userscontroller "github.com/Dwibi/beli-mang/src/http/controllers/users"

func (r *Router) RegisterUser() {
	v1userscontroller := v1userscontroller.New(
		&v1userscontroller.V1Users{
			DB: r.DB,
		})

	r.Router.Post("/admin/register", v1userscontroller.AdminRegister)
	r.Router.Post("/admin/login", v1userscontroller.AdminLogin)
	r.Router.Post("/users/register", v1userscontroller.UserRegister)
	r.Router.Post("/users/login", v1userscontroller.UserLogin)

}
