package address

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-ralali/apiDemo/api/account/address/util"
	"github.com/mayur-ralali/apiDemo/lib"
	"github.com/mayur-tolexo/flash"
)

//Address services
type Address struct {
	flash.Server `prefix:"/account/" v:"1" root:"/address/"`
	listAddress  flash.GET `url:"/"`
}

//listAddress will list address details
func (*Address) ListAddress(c *gin.Context) {

	var (
		data interface{}
		err  error
	)
	if err = util.ValidateListAddress(c); err == nil {
		data, err = util.ProcessListAddress(c)
	}
	lib.BuildResp(c, data, err)
}
