package controllers

import (
	"github.com/pompeu/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func ShowPost(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/post/show/", "", 1)
	post := new(models.Post).GetPost(id)

	tmpl, err := template.ParseFiles("../pompeu/templates/show-posts.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, post)
}

func CriarPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")

	tmpl, err := template.ParseFiles("../pompeu/templates/posts.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

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
			err = tmpl.Execute(w, nil)
		}
	}

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
