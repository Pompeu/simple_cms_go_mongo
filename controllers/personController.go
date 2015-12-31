package controllers

import (
	"github.com/pompeu/models"
	"html/template"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 301)
}

func Login(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)
	w.Header().Add("Content Type", "text/html")

	tmpl, err := template.ParseFiles("../pompeu/templates/login.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		invalid := validImputs("ok", email, password)

		if invalid.Value != "done" {
			err = tmpl.Execute(w, invalid)
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
		err = tmpl.Execute(w, nil)
	}

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func Registrar(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	var invalid Invalid
	tmpl, err := template.ParseFiles("../pompeu/templates/registrar.html")

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		invalid = validImputs(name, email, password)

		if invalid.Value != "done" {
			err = tmpl.Execute(w, invalid)
		} else {
			hash := genereteCode([]byte(password))
			p := &models.Person{}
			if err := p.Create(name, email, hash); err == nil {
				http.Redirect(w, r, "/", 301)
			}
		}
	} else {
		err = tmpl.Execute(w, nil)
	}

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
