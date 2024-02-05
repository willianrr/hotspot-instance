package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Email      string
	Phone      string
	FirstLogin bool
	Firstname  string
	Lastname   string
	Role       string
	Password   string
}

type UserResponse struct {
	gorm.Model
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	FirstLogin bool   `json:"firstlogin"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Role       string `json:"role"`
}
