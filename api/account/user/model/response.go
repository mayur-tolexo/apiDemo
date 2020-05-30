package model

import (
	"github.com/mayur-tolexo/apiDemo/db"
)

type ListUser struct {
	User       []UserDetail `json:"user"`
	TotalCount int          `json:"total_count"`
}

type UserDetail struct {
	db.User //we are using db user struct as embedded struct as all fields are same
}
