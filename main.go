package main

import (
	"github.com/pompeu/routers"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	reHandler := routers.ReHander()
	http.ListenAndServe(":"+port, reHandler)
}
