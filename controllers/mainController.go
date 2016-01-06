package controllers

import (
	"github.com/pompeu/models"
	"net/http"
	"regexp"
)

type Main struct {
	User  string
	Posts []models.Post
}

func MainIndex(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	var tag string
	for _, v := range r.URL.Query() {
		tag = v[0]
	}
	tmpl := TemplateParse("../pompeu/templates/index.html", w)
	main := new(Main)
	main.User = getUserName(r)
	main.Posts = new(models.Post).GetPosts()
	if t, _ := regexp.MatchString("^[a-zA-Z]{2,}$", q); t {
		main.Posts = new(models.Post).FindByName(q)
	}

	if t, _ := regexp.MatchString("^[a-zA-Z]{3,}$", tag); t {
		main.Posts = new(models.Post).GetPostsByTag(tag)
	}
	tmpl.Execute(w, main)
}
