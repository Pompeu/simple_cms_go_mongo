package models

import (
	"encoding/json"
	"fmt"
	"github.com/pompeu/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"github.com/pompeu/db"
	"net/http"
	"strings"
)

type Person struct {
	Id       bson.ObjectId `json:"id" bson:"_id`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

func (p *Person) Create(name, email, hash string) error {
	p.Id = bson.NewObjectId()
	p.Name = name
	p.Email = email
	p.Password = hash
	session := db.SimpleSession("persons")
	defer session.Close()
	err := session.DB("test").C("persons").Insert(p)
	return err
}

func (p *Person) Login(email string) (Person, error) {
	session := db.SimpleSession("persons")
	err := session.DB("test").C("persons").Find(bson.M{"email": email}).One(&p)
	defer session.Close()
	return *p, err
}

func (p *Person) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	person := Person{}
	json.NewDecoder(r.Body).Decode(&person)
	person.Id = bson.NewObjectId()
	db.Session("persons").Insert(person)
	personJson, _ := json.Marshal(person)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", personJson)
}

func (p *Person) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.Replace(r.URL.Path, "/users/", "", 1)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	person := Person{}
	db.Session("persons").Find(bson.M{"id": oid}).One(&person)
	w.WriteHeader(http.StatusOK)
	personJson, _ := json.Marshal(person)
	fmt.Fprintf(w, "%s", personJson)
}

func (p *Person) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []Person
	db.Session("persons").Find(bson.M{}).All(&persons)
	w.WriteHeader(http.StatusOK)
	jsonPersons, _ := json.Marshal(persons)
	fmt.Fprintf(w, "%s", jsonPersons)
}

func (p *Person) Update(w http.ResponseWriter, r *http.Request) {

	id := strings.Replace(r.URL.Path, "/users/", "", 1)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	person := &Person{}
	err := json.NewDecoder(r.Body).Decode(person)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Session("persons").Update(bson.M{"id": oid}, person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (p *Person) Remove(w http.ResponseWriter, r *http.Request) {

	id := strings.Replace(r.URL.Path, "/users/", "", 1)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	db.Session("persons").Remove(bson.M{"id": oid})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
