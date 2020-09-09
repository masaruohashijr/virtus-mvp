package security

import (
	mdl "beerwh/models"
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"
)

var Authenticated = false

func CheckInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	if !Authenticated {
		http.Redirect(w, r, "/login", 301)
		return false
	}
	return true
}

func getLoggedUser(r *http.Request) mdl.User {
	var user mdl.User
	var store = sessions.NewCookieStore([]byte("beerwh"))
	session, _ := store.Get(r, "beerwh")
	sessionUser := session.Values["user"]
	if sessionUser != nil {
		strUser := sessionUser.(string)
		json.Unmarshal([]byte(strUser), &user)
	}
	return user
}
