package controller

import (
	"fmt"
	"net/http"
	"vnia-auth-session/helper"
	"vnia-auth-session/models"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var Store = sessions.NewCookieStore([]byte("vnia-sessions"))

type ControllerUsersImpl struct {
	Db *gorm.DB
}

func NewControllerUsers(db *gorm.DB) *ControllerUsersImpl {
	return &ControllerUsersImpl{
		Db: db,
	}
}

func (c *ControllerUsersImpl) Register(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["auth"] == "true" {
		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}
	Username := r.FormValue("username")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := &models.User{
		Username: Username,
		Email:    Email,
		Password: string(hash),
	}
	if err := c.Db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, user)
}

func (c *ControllerUsersImpl) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["auth"] == "true" {
		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	var user models.User
	if err := c.Db.Where("email = ?", Email).Take(&user).Error; err == gorm.ErrRecordNotFound {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password)); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	session.Save(r, w)
}

func (c *ControllerUsersImpl) Home(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["auth"] == "true" {
		fmt.Fprint(w, "Succes Login")
		return
	}
	http.Redirect(w, r, "/login", http.StatusUnauthorized)
}
