package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-tolexo/apiDemo/api/account/user/model"
	"github.com/mayur-tolexo/apiDemo/api/account/user/util"
	"github.com/mayur-tolexo/apiDemo/lib"
	"github.com/mayur-tolexo/flash"
)

//User services
type User struct {
	flash.Server `prefix:"/account/" v:"1" root:"/user/"`
	listUser     flash.GET  `url:"/"`
	createUser   flash.POST `url:"/"`
}

//ListUser will list user details
func (*User) ListUser(c *gin.Context) {

	var (
		data interface{}
		err  error
	)
	if err = util.ValidateListUser(c); err == nil {
		data, err = util.ProcessListUser(c)
	}
	lib.BuildResp(c, data, err)
}

//CreateUser will create new user
func (*User) CreateUser(c *gin.Context) {

	var (
		req  model.CreateUser
		data interface{}
		err  error
	)
	if req, err = util.ValidateCreateUser(c); err == nil {
		data, err = util.ProcessCreateUser(req)
	}
	lib.BuildResp(c, data, err)
}
