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

	tmpl, err := template.ParseFiles("../pompeu/templates/person.tmpl")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	invalid := &Invalid{"pass", "none"}
	err = tmpl.Execute(w, invalid)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
