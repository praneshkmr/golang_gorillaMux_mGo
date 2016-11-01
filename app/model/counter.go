package model

import "gopkg.in/mgo.v2/bson"

// Counter is the Data Model to store counter related data
type Counter struct {
	_ID      bson.ObjectId `bson:"_id,omitempty"`
	Resource string        `bson:"resource"`
	C        int           `bson:"c"`
}
