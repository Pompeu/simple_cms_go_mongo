package routers

import (
	"github.com/pompeu/controllers"
)

func ReHander() *RegexpHandler {

	h := new(RegexpHandler)

	h.HandleFunc("/post/$", "GET", controllers.CriarPost)
	h.HandleFunc("/post/$", "POST", controllers.CriarPost)

	h.HandleFunc("/post/remove/[0-9a-z]+$", "GET", controllers.RemovePost)
	h.HandleFunc("/post/edit/[0-9a-z]+$", "GET", controllers.EditarPost)
	h.HandleFunc("/post/edit/+$", "POST", controllers.EditarPost)
	h.HandleFunc("/post/show/[0-9a-z]+$", "GET", controllers.ShowPost)

	h.HandleFunc("/tags/[a-zA-Z]+$", "GET", controllers.ShowPostByTags)

	h.HandleFunc("/login/", "GET", controllers.Login)
	h.HandleFunc("/login/", "POST", controllers.Login)
	h.HandleFunc("/logout/", "GET", controllers.Logout)

	h.HandleFunc("/registrar/", "GET", controllers.Registrar)
	h.HandleFunc("/registrar/", "POST", controllers.Registrar)

	h.HandleFunc(".*.[js|css|png|svg|jpg]", "GET", controllers.Assets)

	h.HandleFunc("/", "GET", controllers.MainIndex)

	return h
}
