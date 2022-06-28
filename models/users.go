package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id" sql:"id"`
	Username string `gorm:"not null" json:"username" sql:"usernams"`
	Email    string `gorm:"not null" json:"email" sql:"email"`
	Password string `gorm:"not null" json:"password" sql:"password"`
}
