package security

import (
	"net/http"
)

var Authenticated = false

func CheckInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) {
	if !Authenticated {
		http.Redirect(w, r, "/login", 301)
	}
}
