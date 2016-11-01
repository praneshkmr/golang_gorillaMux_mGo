package model

import "gopkg.in/mgo.v2/bson"

//User is the Data Model to store user related data
type User struct {
	ID       int           `bson:"id" json:"id"`
	_ID      bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
}
