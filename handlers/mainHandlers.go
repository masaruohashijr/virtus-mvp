package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	mdl "virtus/models"
	route "virtus/routes"
	sec "virtus/security"
)

var Db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if sec.IsAuthenticated(w, r) {
		http.Redirect(w, r, route.EntidadesRoute, 200)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
	log.Println("IndexHandler")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout Handler")
	session, _ := sec.Store.Get(r, sec.CookieName)
	delete(session.Values, "user")
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.ServeFile(w, r, "tiles/login.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "tiles/login.html")
		return
	}
	username := r.FormValue("usrname")
	password := r.FormValue("psw")
	var user mdl.User
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Println(string(bytes))
	Db.QueryRow("SELECT id, name, "+
		" username, password, COALESCE(role_id, 0)"+
		" FROM users WHERE username=$1", &username).Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Role)
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("erro do /login")
		http.Redirect(w, r, "/login", 301)
	}

	AddUserInCookie(w, r, user)
	// Abrindo o Cookie
	savedUser := GetUserInCookie(w, r)
	log.Println("MAIN Saved User is " + savedUser.Username)
	http.Redirect(w, r, route.EntidadesRoute, 301)
}

func GetUserInCookie(w http.ResponseWriter, r *http.Request) mdl.User {
	var savedUser mdl.User
	session, _ := sec.Store.Get(r, sec.CookieName)
	sessionUser := session.Values["user"]
	if sessionUser != nil {
		strUser := sessionUser.(string)
		json.Unmarshal([]byte(strUser), &savedUser)
	}
	log.Println("Saved User is " + savedUser.Name)
	return savedUser
}

func AddUserInCookie(w http.ResponseWriter, r *http.Request, user mdl.User) {
	sec.Store.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session, _ := sec.Store.Get(r, sec.CookieName)
	bytesUser, _ := json.Marshal(&user)
	session.Values["user"] = string(bytesUser)
	sec.Store.Save(r, w, session)
	session.Save(r, w)
}

func BuildLoggedUser(user mdl.User) mdl.LoggedUser {
	var loggedUser mdl.LoggedUser
	loggedUser.User = user
	loggedUser.HasPermission = func(permission string) bool {
		//log.Println(permission)
		query := "SELECT " +
			"A.feature_id, B.code FROM features_roles A, features B " +
			"WHERE A.feature_id = B.id AND A.role_id = $1"
		rows, _ := Db.Query(query, user.Role)
		var features []mdl.Feature
		var feature mdl.Feature
		for rows.Next() {
			rows.Scan(&feature.Id, &feature.Code)
			features = append(features, feature)
		}
		for _, value := range features {
			//log.Println(value.Code)
			if value.Code == permission {
				//log.Println(permission + " encontrada!!!")
				return true
			}
		}
		return false
	}
	return loggedUser
}
