package entity

import "time"

type User struct {
	Id             int
	UserName       string    `json:"user_name"`
	Password       string    `json:"password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Role           string    `json:"role"`
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	ProfilePicture string    `json:"profile_picture"`
	Blood          string    `json:"blood"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
