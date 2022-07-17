package formatter

import (
	"aplikasi_golang/entity"

	"github.com/ekokurniadi/indodate"
)

type UserFormatter struct {
	Id             int    `json:"id"`
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Role           string `json:"role"`
	Age            int    `json:"age"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
	Blood          string `json:"blood"`
	Keterangan     string `json:"keterangan"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func FormatSingleUser(user entity.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.Id = user.Id
	userFormatter.UserName = user.UserName
	userFormatter.FirstName = user.FirstName
	userFormatter.LastName = user.LastName
	userFormatter.Role = user.Role
	var keteranganUser string
	if user.Age < 30 {
		keteranganUser = "Belum Dewasa"
	} else {
		keteranganUser = "Sudah Dewasa"
	}
	userFormatter.Keterangan = keteranganUser
	userFormatter.Age = user.Age
	userFormatter.Gender = user.Gender
	userFormatter.Address = user.Address
	userFormatter.ProfilePicture = user.ProfilePicture
	userFormatter.CreatedAt = indodate.DateWithMin(user.CreatedAt)
	userFormatter.UpdatedAt = indodate.DateWithMin(user.UpdatedAt)

	return userFormatter
}

func FormatManyUser(users []entity.User) []UserFormatter {
	var userFormatter []UserFormatter
	for _, user := range users {
		userFormat := FormatSingleUser(user)
		userFormatter = append(userFormatter, userFormat)
	}
	return userFormatter
}
