package main

import (
	"github.com/mayur-ralali/apiDemo/api/account/user"
	"github.com/mayur-ralali/apiDemo/api/account/address"
	"github.com/mayur-tolexo/flash"
)

func main() {
	router := flash.Default()
	router.AddService(&user.User{})
	router.AddService(&address.Address{})
	router.Start(":3030")
}
