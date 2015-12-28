package routers

import (
	"github.com/pompeu/helpers"
	"github.com/pompeu/models"
)

func ReHander() *helpers.RegexpHandler {
	person := &models.Person{}
	post := &models.Post{}
	coment := &models.Coment{}
	reHandler := new(helpers.RegexpHandler)

	reHandler.HandleFunc("/users/$", "POST", person.Save)
	reHandler.HandleFunc("/users/$", "GET", person.GetAll)
	reHandler.HandleFunc("/users/[0-9a-z]+$", "PUT", person.Update)
	reHandler.HandleFunc("/users/[0-9a-z]+$", "DELETE", person.Remove)
	reHandler.HandleFunc("/users/[0-9a-z]+$", "GET", person.GetOne)

	reHandler.HandleFunc("/posts/$", "POST", post.Save)
	reHandler.HandleFunc("/posts/$", "GET", post.GetAll)
	reHandler.HandleFunc("/posts/[0-9a-z]+$", "PUT", post.Update)
	reHandler.HandleFunc("/posts/[0-9a-z]+$", "DELETE", post.Remove)
	reHandler.HandleFunc("/posts/[0-9a-z]+$", "GET", post.GetOne)

	reHandler.HandleFunc("/coments/$", "POST", coment.Save)
	reHandler.HandleFunc("/coments/$", "GET", coment.GetAll)
	reHandler.HandleFunc("/coments/[0-9a-z]+$", "PUT", coment.Update)
	reHandler.HandleFunc("/coments/[0-9a-z]+$", "DELETE", coment.Remove)
	reHandler.HandleFunc("/coments/[0-9a-z]+$", "GET", coment.GetOne)

	return reHandler
}
