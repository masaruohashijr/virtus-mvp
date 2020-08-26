package handlers

import (
	"database/sql"
	mdl "diaria/models"
	route "diaria/routes"
	sec "diaria/security"
	"golang.org/x/crypto/bcrypt"
	//	"log"
	"net/http"
	//	"strconv"
)

var Db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//	log.Println("sec.Authenticated: " + strconv.FormatBool(sec.Authenticated))
	sec.IsAuthenticated(w, r)
	http.Redirect(w, r, route.MealsRoute, 200)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "tmpl/login.html")
		return
	}
	username := r.FormValue("usrname")
	password := r.FormValue("psw")
	var user mdl.User
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := Db.QueryRow("SELECT username, password FROM users WHERE username=$1", &username).Scan(&user.Username, &user.Password)
	sec.CheckInternalServerError(err, w)
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
	}
	sec.Authenticated = true
	http.Redirect(w, r, route.MealsRoute, 301)
}
