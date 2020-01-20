package db

import "time"

type User struct {
	UserID    int       `json:"user_id" sql:"user_id"`
	Name      string    `json:"name" sql:"name"`
	Mobile    string    `json:"mobile" sql:"mobile"`
	Email     string    `json:"email" sql:"email"`
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	CreatedBy int       `json:"created_by" sql:"created_by"`
}
