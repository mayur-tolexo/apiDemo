package main

import (
	"github.com/mayur-ralali/apiDemo/api/account"
	"github.com/mayur-ralali/apiDemo/lib"
	"github.com/mayur-tolexo/flash"
)

func main() {
	router := flash.Default()
	lib.AddServices(&router, account.GetServices()...)
	router.Start(":3030")
}
