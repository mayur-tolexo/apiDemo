package model

import (
	"github.com/mayur-tolexo/apiDemo/db"
)

type ListAddress struct {
	Address    []AddressDetail `json:"addresss"`
	TotalCount int             `json:"total_count"`
}

//AddressDetail deatils
type AddressDetail struct {
	db.Address //we are using db user struct as embedded struct as all fields are same
}
