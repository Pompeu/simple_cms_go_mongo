package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type Invalid struct {
	Input string
	Value string
}

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

	tmpl, err := template.ParseFiles("../pompeu/templates/registrar.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	title := &Server{"Registrar"}
	err = tmpl.Execute(w, title)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
