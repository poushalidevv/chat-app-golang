package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserLogin struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}