package mongo

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func getMongoSession() *mgo.Session {
	if session == nil {
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		return session
	}
	return session.Copy()
}
