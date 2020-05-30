package util

import (
	"github.com/gin-gonic/gin"
	"github.com/mayur-tolexo/apiDemo/api/account/address/model"
)

//ProcessListAddress will process the list address request
func ProcessListAddress(c *gin.Context) (data model.ListAddress, err error) {
	data.Address = make([]model.AddressDetail, 0)
	//TODO: create database connection and fetch address details
	return
}
