package controllers

import (
	"html/template"
	"net/http"
)

func TemplateParse(path string, w http.ResponseWriter) (tmpl *template.Template) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := tmpl.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	return tmpl
}
