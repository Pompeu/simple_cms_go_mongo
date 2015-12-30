package main

import (
	"github.com/pompeu/routers"
	"log"
	"net/http"
)

func main() {
	log.Print("server on")
	reHandler := routers.ReHander()
	http.ListenAndServe(":3000", reHandler)
}
