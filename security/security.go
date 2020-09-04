package security

import (
	mdl "beerwh/models"
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
