package models

import (
	"github.com/jinzhu/gorm"
)

// User model definition
type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"type:varchar(255);unique_index"`
	Email    string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255)"`
}

type UserLogin struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}