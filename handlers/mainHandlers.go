package handlers

import (
	mdl "beerwh/models"
	route "beerwh/routes"
	sec "beerwh/security"
	"database/sql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var Db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//sec.IsAuthenticated(w, r)
	http.Redirect(w, r, route.BeersRoute, 200)
}

var store = sessions.NewCookieStore([]byte("beerwh"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("entrei login")
		http.ServeFile(w, r, "tmpl/login.html")
		return
	}
	username := r.FormValue("usrname")
	log.Println("user")
	password := r.FormValue("psw")
	log.Println("senha")
	var user mdl.User
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := Db.QueryRow("SELECT id, username, password FROM clients WHERE username=$1", &username).Scan(&user.Id, &user.Username, &user.Password)
	sec.CheckInternalServerError(err, w)
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("erro do /login")
		http.Redirect(w, r, "/login", 301)
	}
	sec.Authenticated = true
	session, _ := store.Get(r, "beerwh")
	bytesUser, _ := json.Marshal(user)
	session.Values["user"] = string(bytesUser)
	sessions.Save(r, rw)
	log.Println("antes OrdersRoute")
	http.Redirect(w, r, route.OrdersRoute, 301)
}
