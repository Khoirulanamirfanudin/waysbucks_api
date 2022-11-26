package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"emai"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
