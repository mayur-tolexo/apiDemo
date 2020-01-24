package main

import (
	"github.com/mayur-ralali/apiDemo/api/account"
)

func main() {
	router := account.SetupRouter()
	router.Start(":3030")
}
