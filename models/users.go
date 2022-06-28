package models

import (
	"github.com/jinzhu/gorm"
)

type SignupForm struct {
	Username string `scheme:"username"`
	Email    string `scheme:"email"`
	Password string `scheme:"password"`
}

type SigninForm struct {
	Email    string `scheme:"email"`
	Password string `scheme:"password"`
}

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id" sql:"id"`
	Username string `gorm:"not null" json:"username" sql:"usernams"`
	Email    string `gorm:"not null" json:"email" sql:"email"`
	Password string `gorm:"not null" json:"password" sql:"password"`
}
