package mongo

import (
	"./../../../app/model"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var host string
var database string

// GetMongoSession returns the MGo's Session Object upon Request
func GetMongoSession() *mgo.Session {
	if session == nil {
		session, err := mgo.Dial(host)
		if err != nil {
			panic(err)
		}
		return session
	}
	return session.Copy()
}

func GetDatabase() string {
	return database
}

func SetConfig(config *model.Configuration) {
	host = config.MongoDB.Host
	database = config.MongoDB.Database
}
