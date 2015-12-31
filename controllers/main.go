package controllers

import (
	"github.com/pompeu/models"
	"html/template"
	"log"
	"net/http"
)

type Main struct {
	User  string
	Posts []models.Post
}

func MainIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")

	tmpl, err := template.ParseFiles("../pompeu/templates/index.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	main := new(Main)
	main.User = getUserName(r)
	main.Posts = new(models.Post).GetPosts()
	err = tmpl.Execute(w, main)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
