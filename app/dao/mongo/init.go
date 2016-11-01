package mongo

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetMongoSession returns the MGo's Session Object upon Request
func GetMongoSession() *mgo.Session {
	if session == nil {
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		return session
	}
	return session.Copy()
}
