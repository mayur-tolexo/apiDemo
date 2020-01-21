package lib

import (
	"fmt"
	"github.com/mayur-tolexo/flash"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	services = make(map[string]map[string]interface{})
	mux      sync.Mutex
)

//BuildResp will build response model
func BuildResp(c *gin.Context, data interface{}, err error) {
	httpCode := getStatusCode(err)
	resp := gin.H{
		"data": data,
	}
	if err != nil {
		resp["error"] = err
	}
	c.JSON(httpCode, resp)
}

//AddService will add services
func AddService(module string, service interface{}) {
	ref := reflect.TypeOf(service)
	if ref.Kind() == reflect.Ptr {
		path := ref.Elem().PkgPath()
		name := reflect.ValueOf(service).Type().String()
		key := path + ":" + name
		fmt.Println(key)
		mux.Lock()
		defer mux.Unlock()
		if services[module] == nil {
			services[module] = make(map[string]interface{})
		}
		services[module][key] = service
	}
}

//AddServices will add services in router
func AddServices(router *flash.Server, services ...interface{}) {
	for _, s := range services {
		router.AddService(s)
	}
}

//GetServices will return all services
func GetServices(module string) (all []interface{}) {
	mux.Lock()
	defer mux.Unlock()
	for _, s := range services[module] {
		all = append(all, s)
	}
	return
}

func getStatusCode(err error) (code int) {
	//TODO: compare error to provide accurate status code
	if err != nil {
		code = 500
	} else {
		code = 200
	}
	return
}
