package controller

import (
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
	if session.Values["auth-vnia"] == "true" {
		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}
	Username := r.FormValue("username")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
	  helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user := &models.User{
		Username: Username,
		Email:    Email,
		Password: string(hash),
	}
	if err := c.Db.Create(&user).Error; err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondwithJSON(w, http.StatusOK, user)
}

func (c *ControllerUsersImpl) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["auth-vnia"] == "true" {
		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	var user models.User
	if err := c.Db.Where("email = ?", Email).Take(&user).Error; err == gorm.ErrRecordNotFound {
		helper.RespondWithError(w, http.StatusInternalServerError, "Akun Kamu Belum Terdaftar")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password)); err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	session.Values["auth-vnia"] = true
	session.Values["email"] = Email
	session.Values["username"] = user.Username
	session.Save(r, w)
	helper.RespondwithJSON(w, http.StatusOK, user)
}

func (c *ControllerUsersImpl) Home(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["auth-vnia"] == true {
	  var user models.User
	  c.Db.Where("email = ?", session.Values["email"]).Take(&user)
	  helper.RespondwithJSON(w, http.StatusOK, user)
		return
	} else {
	  helper.RespondWithError(w, http.StatusUnauthorized, "Silahkan Login")
	}
}
