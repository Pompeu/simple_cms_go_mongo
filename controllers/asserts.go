package controllers

import (
	"net/http"
	"strings"
)

func Assets(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if suffixTrusted(path) {
		http.ServeFile(w, r, r.URL.Path[1:])
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	}
}

func suffixTrusted(path string) bool {
	return strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".js")
}
