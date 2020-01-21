package account

//blank import to trigger inits
import (
	_ "github.com/mayur-ralali/apiDemo/api/account/address"
	_ "github.com/mayur-ralali/apiDemo/api/account/user"
	"github.com/mayur-ralali/apiDemo/constant"
	"github.com/mayur-ralali/apiDemo/lib"
)

//GetServices will trigger all account submodule init and return services
func GetServices() []interface{} {
	return lib.GetServices(constant.Account)
}
