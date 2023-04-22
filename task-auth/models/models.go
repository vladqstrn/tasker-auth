package models

import "time"

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Created_at time.Time
}
