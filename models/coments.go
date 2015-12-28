package models

import (
	"encoding/json"
	"fmt"
	"github.com/pompeu/db"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type Coment struct {
	Id    bson.ObjectId `json:"id" bson:"_id`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"body" bson:"body"`
}

func (c *Coment) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	coment := Coment{}
	json.NewDecoder(r.Body).Decode(&coment)
	coment.Id = bson.NewObjectId()
	db.Session("coments").Insert(coment)
	comentJson, _ := json.Marshal(coment)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", comentJson)
}

func (c *Coment) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.Replace(r.URL.Path, "/users/", "", 1)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	coment := Coment{}
	db.Session("coments").Find(bson.M{"id": oid}).One(&coment)
	w.WriteHeader(http.StatusOK)
	comentJson, _ := json.Marshal(coment)
	fmt.Fprintf(w, "%s", comentJson)
}

func (c *Coment) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var coments []Coment
	db.Session("coments").Find(bson.M{}).All(&coments)
	w.WriteHeader(http.StatusOK)
	jsonComents, _ := json.Marshal(coments)
	fmt.Fprintf(w, "%s", jsonComents)
}

func (c *Coment) Update(w http.ResponseWriter, r *http.Request) {

	id := strings.Replace(r.URL.Path, "/users/", "", 1)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	coment := &Coment{}
	err := json.NewDecoder(r.Body).Decode(coment)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Session("coments").Update(bson.M{"id": oid}, coment)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (c *Coment) Remove(w http.ResponseWriter, r *http.Request) {

	id := strings.Replace(r.URL.Path, "/users/", "", 1)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	db.Session("coments").Remove(bson.M{"id": oid})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
