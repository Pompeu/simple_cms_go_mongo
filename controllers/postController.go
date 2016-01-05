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

type Post struct {
	Id    string `json:"id" bson:"id`
	Title string `json:"title" bson:"title"`
	Body  string `json:"body" bson:"body"`
	Tags  string `json:"tags" bson:"tags"`
}

func EditarPost(w http.ResponseWriter, r *http.Request) {
	tmpl := TemplateParse("../pompeu/templates/edit-posts.html", w)
	middlware(w, r)

	if r.Method == "POST" {
		id := r.FormValue("id")
		title := r.FormValue("title")
		body := r.FormValue("body")
		tags := strings.Split(r.FormValue("tags"), " ")
		err := new(models.Post).Update(id, title, body, tags)
		if err != nil {
			fmt.Println(err)
		} else {
			http.Redirect(w, r, "/", 301)
		}
	} else {
		if id := strings.Replace(
			r.URL.Path, "/post/edit/", "", 1); id != "/post/" {
			post := new(models.Post).GetPost(id)
			tags := strings.Join(post.Tags, " ")
			p := &Post{post.Id.Hex(), post.Title, post.Body, tags}
			tmpl.Execute(w, p)
		} else {
			tmpl.Execute(w, nil)
		}
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
	}
	tmpl.Execute(w, nil)
}
