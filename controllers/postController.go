package controllers

import (
	"github.com/pompeu/models"
	"net/http"
	"strings"
)

func ShowPost(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/post/show/", "", 1)
	post := new(models.Post).GetPost(id)
	TemplateParse("../pompeu/templates/show-posts.html", w).Execute(w, post)
}

func CriarPost(w http.ResponseWriter, r *http.Request) {
	tmpl := TemplateParse("../pompeu/templates/posts.html", w)

	if r.Method == "POST" {
		title := r.FormValue("title")
		body := r.FormValue("body")
		tags := r.FormValue("tags")
		post := &models.Post{}

		post.Title = title
		post.Body = strings.Trim(body, " ")
		post.Tags = strings.Split(tags, " ")
		if _, err := post.Create(); err == nil {
			http.Redirect(w, r, "/", 301)
		} else {
			tmpl.Execute(w, err)
		}
	} else {
		if id := strings.Replace(r.URL.Path, "/post/edit/", "", 1); id != "/post/" {
			post := new(models.Post).GetPost(id)
			tmpl.Execute(w, post)
		} else {
			tmpl.Execute(w, nil)
		}
	}

}
