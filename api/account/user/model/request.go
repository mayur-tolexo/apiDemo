package model

//CreateUser contains the request format of create user api
type CreateUser struct {
	Name   string `json:"name" validate:"required"`
	Mobile string `json:"mobile" validate:"required"`
	Email  string `json:"email" validate:"email"`
}
