package routers

import (
	"github.com/pompeu/controllers"
	"github.com/pompeu/helpers"
	"github.com/pompeu/models"
)

func ReHander() *helpers.RegexpHandler {
	person := &models.Person{}
	post := &models.Post{}
	coment := &models.Coment{}
	server := &controllers.Server{}

	h := new(helpers.RegexpHandler)

	h.HandleFunc("/users/$", "POST", person.Save)
	h.HandleFunc("/users/$", "GET", person.GetAll)
	h.HandleFunc("/users/[0-9a-z]+$", "PUT", person.Update)
	h.HandleFunc("/users/[0-9a-z]+$", "DELETE", person.Remove)
	h.HandleFunc("/users/[0-9a-z]+$", "GET", person.GetOne)

	h.HandleFunc("/posts/$", "POST", post.Save)
	h.HandleFunc("/posts/$", "GET", post.GetAll)
	h.HandleFunc("/posts/[0-9a-z]+$", "PUT", post.Update)
	h.HandleFunc("/posts/[0-9a-z]+$", "DELETE", post.Remove)
	h.HandleFunc("/posts/[0-9a-z]+$", "GET", post.GetOne)

	h.HandleFunc("/coments/$", "POST", coment.Save)
	h.HandleFunc("/coments/$", "GET", coment.GetAll)
	h.HandleFunc("/coments/[0-9a-z]+$", "PUT", coment.Update)
	h.HandleFunc("/coments/[0-9a-z]+$", "DELETE", coment.Remove)
	h.HandleFunc("/coments/[0-9a-z]+$", "GET", coment.GetOne)

	h.HandleFunc("/login/", "GET", controllers.Login)
	h.HandleFunc("/login/", "POST", controllers.Login)

	h.HandleFunc("/registrar/", "GET", controllers.Registrar)
	h.HandleFunc("/registrar/", "POST", controllers.Registrar)

	h.HandleFunc(".*.[js|css|png|svg|jpg]", "GET", controllers.Assets)

	h.HandleFunc("/", "GET", server.MainIndex)

	return h
}
