package controllers

import (
	"github.com/pompeu/Godeps/_workspace/src/github.com/gorilla/securecookie"
	"net/http"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

const session = "sessionName"

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie(session); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode(session, cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode(session, value); err == nil {
		cookie := &http.Cookie{
			Name:  session,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}
func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   session,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
