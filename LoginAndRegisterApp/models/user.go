package models

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"time created"`
}
