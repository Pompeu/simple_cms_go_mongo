package routers

import (
	"github.com/pompeu/controllers"
	"github.com/pompeu/helpers"
)

func ReHander() *helpers.RegexpHandler {

	h := new(helpers.RegexpHandler)

	h.HandleFunc("/post/$", "GET", controllers.CriarPost)
	h.HandleFunc("/post/edit/[0-9a-z]+$", "GET", controllers.CriarPost)
	h.HandleFunc("/post/show/[0-9a-z]+$", "GET", controllers.ShowPost)
	h.HandleFunc("/post/", "POST", controllers.CriarPost)

	h.HandleFunc("/login/", "GET", controllers.Login)
	h.HandleFunc("/login/", "POST", controllers.Login)
	h.HandleFunc("/logout/", "GET", controllers.Logout)

	h.HandleFunc("/registrar/", "GET", controllers.Registrar)
	h.HandleFunc("/registrar/", "POST", controllers.Registrar)

	h.HandleFunc(".*.[js|css|png|svg|jpg]", "GET", controllers.Assets)

	h.HandleFunc("/", "GET", controllers.MainIndex)

	return h
}
