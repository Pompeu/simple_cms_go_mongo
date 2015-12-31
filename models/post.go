package models

import (
	"encoding/json"
	"fmt"
	"github.com/pompeu/db"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type Post struct {
	Id    bson.ObjectId `json:"id" bson:"_id`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"body" bson:"body"`
	Tags  []string      `json:"tags" bson:"tags"`
}

func (p *Post) Create() (Post, error) {
	p.Id = bson.NewObjectId()
	session := db.SimpleSession("posts")
	err := session.DB("test").C("posts").Insert(p)
	defer session.Close()
	return *p, err
}

func (p *Post) GetPosts() []Post {
	var posts []Post
	session := db.SimpleSession("posts")
	err := session.DB("test").C("posts").Find(bson.M{}).All(&posts)
	fmt.Println(err)
	defer session.Close()
	return posts
}

func (p *Post) GetPost(id string) Post {
	session := db.SimpleSession("posts")
	oid := bson.ObjectIdHex(id)
	err := session.DB("test").C("posts").Find(bson.M{"id": oid}).One(&p)
	fmt.Println(err)
	defer session.Close()
	return *p
}

func (p *Post) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := Post{}
	json.NewDecoder(r.Body).Decode(&post)
	post.Id = bson.NewObjectId()
	db.Session("posts").Insert(post)
	postJson, _ := json.Marshal(post)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", postJson)
}

func (p *Post) GetOne(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/posts/", "", 1)
	w.Header().Set("Content-Type", "application/json")
	post := Post{}
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	db.Session("posts").Find(bson.M{"id": oid}).One(&post)
	w.WriteHeader(http.StatusOK)
	postJson, _ := json.Marshal(post)
	fmt.Fprintf(w, "%s", postJson)
}

func (p *Post) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	db.Session("posts").Find(bson.M{}).All(&posts)
	w.WriteHeader(http.StatusOK)
	jsonPosts, _ := json.Marshal(posts)
	fmt.Fprintf(w, "%s", jsonPosts)
}

func (p *Post) Update(w http.ResponseWriter, r *http.Request) {

	id := strings.Replace(r.URL.Path, "/posts/", "", 1)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	post := &Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Session("posts").Update(bson.M{"id": oid}, post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (p *Post) Remove(w http.ResponseWriter, r *http.Request) {

	id := strings.Replace(r.URL.Path, "/posts/", "", 1)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	db.Session("posts").Remove(bson.M{"id": oid})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
