package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-tolexo/apiDemo/api/account/user/model"
)

//ValidateListUser will validate the list user request
func ValidateListUser(c *gin.Context) (err error) {
	return
}

//ValidateListUser will validate the list user request
func ValidateCreateUser(c *gin.Context) (user model.CreateUser, err error) {
	err = c.BindJSON(&user)
	return
}
