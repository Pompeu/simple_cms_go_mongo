package controllers

import (
	"github.com/pompeu/models"
	"net/http"
)

type Main struct {
	User  string
	Posts []models.Post
}

func MainIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := TemplateParse("../pompeu/templates/index.html", w)
	main := new(Main)
	main.User = getUserName(r)
	main.Posts = new(models.Post).GetPosts()
	tmpl.Execute(w, main)
}
