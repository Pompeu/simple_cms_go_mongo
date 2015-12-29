package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	Title string
}

func (t *Server) MainIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content Type", "text/html")

	tmpl, err := template.ParseFiles("../pompeu/templates/index.tmpl")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	title := &Server{"CMS Limp"}
	err = tmpl.Execute(w, title)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
