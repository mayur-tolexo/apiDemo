package db

import "time"

type Address struct {
	AddressID int       `json:"address_id" sql:"address_id"`
	Street    string    `json:"street" sql:"street"`
	City      string    `json:"city" sql:"city"`
	State     string    `json:"state" sql:"state"`
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	CreatedBy int       `json:"created_by" sql:"created_by"`
}
