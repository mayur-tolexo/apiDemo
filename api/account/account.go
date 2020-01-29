package account

import (
	"github.com/mayur-ralali/apiDemo/api/account/address"
	"github.com/mayur-ralali/apiDemo/api/account/user"
	"github.com/mayur-tolexo/flash"
)

//SetupRouter will setup account handlers
func SetupRouter() *flash.Server {
	router := flash.Default()
	router.AddService(&user.User{})
	router.AddService(&address.Address{})
	return router
}
