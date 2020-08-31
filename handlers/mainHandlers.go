package handlers

import (
	mdl "beerwh/models"
	route "beerwh/routes"
	sec "beerwh/security"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var Db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	http.Redirect(w, r, route.OrdersRoute, 200)
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
	err := Db.QueryRow("SELECT id, username, password FROM clients WHERE username=$1", &username).Scan(&user.Id, &user.Username, &user.Password)
	sec.CheckInternalServerError(err, w)
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/login", 301)
	}
	sec.Authenticated = true
	sec.LoggedUser = user
	http.Redirect(w, r, route.OrdersRoute, 301)
}
