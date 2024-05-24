package entities

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

type RegisterParams struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginParams struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}
