package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"-" form:"password" validate:"required"`
	Image    string `json:"image"`
}

type DeleteUserResponse struct {
	ID int `json:"id"`
}
