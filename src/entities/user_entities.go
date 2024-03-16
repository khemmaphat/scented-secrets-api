package entities

import "time"

type User struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	Description string    `json:"description"`
	Telephone   string    `json:"telephone"`
}
