package db

import (
	"errors"
	"gopkg.in/mgo.v2"
)

type MongoDb struct {
	mongoUrl string
	database string
	session  *mgo.Session
}

func NewMongoConnect(mongoUrl, database string) (m MongoDb, e error) {
	m.mongoUrl = mongoUrl
	m.database = database
	session, err := mgo.Dial(m.mongoUrl)
	if err != nil {
		return m, mongoerr(err.Error())
	}
	if err := session.Ping(); err != nil {
		return m, mongoerr(err.Error())
	}
	index := setIndex([]string{"email"})
	if err := session.DB(m.database).C("test").EnsureIndex(index); err != nil {
		return m, mongoerr(err.Error())
	}
	m.session = session
	return
}

func setIndex(key []string) mgo.Index {
	return mgo.Index{
		Key:    key,
		Unique: true,
	}
}

func mongoerr(msg string) error {
	return errors.New("mongodb: " + msg)
}

func (b MongoDb) Close() {
	if b.session != nil {
		b.session.Close()
	}
}

func Session(col string) *mgo.Collection {
	session, err := mgo.Dial("mongodb://localhost/" + col)
	if err != nil {
		panic(err)
	}
	return session.DB("test").C(col)
}

func SimpleSession(col string) *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost/" + col)
	if err != nil {
		panic(err)
	}
	return session
}
