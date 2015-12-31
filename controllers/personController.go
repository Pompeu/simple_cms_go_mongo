package controllers

import (
	"github.com/pompeu/models"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 301)
}

func Login(w http.ResponseWriter, r *http.Request) {

	tmpl := TemplateParse("../pompeu/templates/login.html", w)

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		invalid := validImputs("ok", email, password)

		if invalid.Value != "done" {
			tmpl.Execute(w, invalid)
		} else {
			p := &models.Person{}
			if auth, err := p.Login(email); err == nil {
				valid := compare([]byte(auth.Password), []byte(password))
				if valid == nil {
					setSession(auth.Name, w)
					http.Redirect(w, r, "/", 301)
				} else {
					tmpl.Execute(w, "email or password invalid")
				}
			}
		}
	} else {
		tmpl.Execute(w, nil)
	}

}

func Registrar(w http.ResponseWriter, r *http.Request) {
	tmpl := TemplateParse("../pompeu/templates/registrar.html", w)

	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		invalid := validImputs(name, email, password)

		if invalid.Value != "done" {
			tmpl.Execute(w, invalid)
		} else {
			hash := genereteCode([]byte(password))
			p := &models.Person{}
			if err := p.Create(name, email, hash); err == nil {
				http.Redirect(w, r, "/", 301)
			}
		}
	} else {
		tmpl.Execute(w, nil)
	}
}
