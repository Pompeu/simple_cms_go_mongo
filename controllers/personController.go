package controllers

import (
	"github.com/pompeu/db"
	"github.com/pompeu/models"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content Type", "text/html")

	tmpl, err := template.ParseFiles("../pompeu/templates/login.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	title := &Server{"Login"}
	err = tmpl.Execute(w, title)

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
		log.Println(r.Method)
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		invalid = validImputs(name, email, password)

		if invalid.Value != "done" {
			err = tmpl.Execute(w, invalid)
		} else {
			hash := genereteCode([]byte(password))
			person := &models.Person{bson.NewObjectId(), name, email, hash}
			err := db.Session("persons").Insert(person)

			if err == nil {
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
