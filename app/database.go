package app

import (
	"log"
	"vnia-auth-session/helper"
	"vnia-auth-session/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/restful?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	helper.IfError(err)

	db.AutoMigrate(&models.User{})
	log.Println("Database Connect")
	return db
}
