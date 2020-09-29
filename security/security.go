package security

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var CookieName = "virtus"
var Store = sessions.NewCookieStore([]byte("vindixit123581321"))

func CheckInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	session, err := Store.Get(r, CookieName)
	if err != nil {
		return false
	}
	sessionUser := session.Values["user"]
	if sessionUser == nil {
		return false
	}
	return true
}
