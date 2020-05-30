package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-tolexo/apiDemo/api/account/user/model"
)

//ProcessListUser will process the list user request
func ProcessListUser(c *gin.Context) (data model.ListUser, err error) {
	data.User = make([]model.UserDetail, 0)
	//TODO: create database connection and fetch user details
	return
}

//ProcessCreateUser will process the create user request
func ProcessCreateUser(req model.CreateUser) (UserID int, err error) {
	//TODO: create database connection and create new user
	return
}
