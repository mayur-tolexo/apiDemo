package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-ralali/apiDemo/api/account/user/model"
)

//ProcessListUser will process the list user request
func ProcessListUser(c *gin.Context) (data model.ListUser, err error) {
	return
}

//ProcessCreateUser will process the create user request
func ProcessCreateUser(req model.CreateUser) (UserID int, err error) {
	return
}
