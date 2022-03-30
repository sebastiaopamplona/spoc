package database

import "time"

type User struct {
	UserId    int       `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"createdOn"`
	LastLogin time.Time `json:"lastLogin"`
}
