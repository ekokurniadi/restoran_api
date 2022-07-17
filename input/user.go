package input

type UserInputID struct {
	Id int `uri:"id"`
}

type UserInput struct {
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
}
