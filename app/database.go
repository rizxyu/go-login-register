package app

import (
  "fmt"
	"log"
	"os"
	"vnia-auth-session/helper"
	"vnia-auth-session/models"

  "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
  if err := godotenv.Load(); err != nil {
    panic("Gagal Load Env")
  }
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)), &gorm.Config{})
	helper.IfError(err)

	db.AutoMigrate(&models.User{})
	log.Println("Database Connect")
	return db
}
