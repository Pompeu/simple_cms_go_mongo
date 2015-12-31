package controllers

import (
	"fmt"
	"github.com/pompeu/models"
	"net/http"
	"strings"
)

func ShowPostByTags(w http.ResponseWriter, r *http.Request) {
	tags := strings.Split(r.URL.Path, "/")
	tag := tags[len(tags)-1]
	tmpl := TemplateParse("../pompeu/templates/index.html", w)
	main := new(Main)
	main.User = getUserName(r)
	main.Posts = new(models.Post).GetPostsByTag(tag)
	tmpl.Execute(w, main)
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/post/show/", "", 1)
	post := new(models.Post).GetPost(id)
	tmpl := TemplateParse("../pompeu/templates/show-posts.html", w)
	tmpl.Execute(w, post)
}

func RemovePost(w http.ResponseWriter, r *http.Request) {
	middlware(w, r)
	id := strings.Replace(r.URL.Path, "/post/remove/", "", 1)
	ok, err := new(models.Post).RemovePost(id)
	if ok && err == nil {
		http.Redirect(w, r, "/", 301)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", `<h1> Fail go bottuon back </h1>`)
	}

}
func middlware(w http.ResponseWriter, r *http.Request) {
	if getUserName(r) == "" {
		http.Redirect(w, r, "/", 301)
	} else {
		return
	}
}
func CriarPost(w http.ResponseWriter, r *http.Request) {
	tmpl := TemplateParse("../pompeu/templates/posts.html", w)
	middlware(w, r)

	if r.Method == "POST" {
		title := r.FormValue("title")
		body := r.FormValue("body")
		tags := r.FormValue("tags")

		valid := validPost(title, "ok", tags)

		if valid.Value != "done" {
			tmpl.Execute(w, nil)
		} else {
			post := &models.Post{}
			post.Title = title
			post.Body = strings.Trim(body, " ")
			post.Tags = strings.Split(tags, " ")

			fmt.Println(post)
			if _, err := post.Create(); err == nil {
				http.Redirect(w, r, "/", 301)
			} else {
				tmpl.Execute(w, err)
			}
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
