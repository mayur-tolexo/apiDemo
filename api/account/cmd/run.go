package main

import (
	"github.com/mayur-tolexo/apiDemo/api/account"
)

func main() {
	router := account.SetupRouter()
	router.Start(":3030")
}
