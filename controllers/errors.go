package controllers

import (
	"log"
	"net/http"
)

func CtrlError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
