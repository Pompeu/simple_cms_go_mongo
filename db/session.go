package db

import (
	"gopkg.in/mgo.v2"
)

func Session(col string) *mgo.Collection {
	session, err := mgo.Dial("mongodb://localhost/" + col)
	if err != nil {
		panic(err)
	}
	return session.DB("test").C(col)
}
