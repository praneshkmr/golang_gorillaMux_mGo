package user

import (
	"./../../../model"
	"./../../mongo"
	"gopkg.in/mgo.v2/bson"
)

var database = "mGoMuxRest"
var collection = "users"

//AddUser is the function to add new User
func AddUser(user *model.User) (*model.User, error) {
	session := mongo.GetMongoSession()
	defer session.Close()
	userCollection := session.DB(database).C(collection)
	err := userCollection.Insert(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//GetUser is the function to retrieve User by Id
func GetUser(id int) (*model.User, error) {
	session := mongo.GetMongoSession()
	defer session.Close()
	var user model.User
	userCollection := session.DB(database).C(collection)
	err := userCollection.Find(bson.M{"id": id}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser is the function to update User data by Id
func UpdateUser(id int, user model.User) (*model.User, error) {
	session := mongo.GetMongoSession()
	defer session.Close()
	userCollection := session.DB(database).C(collection)
	query := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email, "password": user.Password}}
	err := userCollection.Update(query, update)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//DeleteUser is the function to delete User by Id
func DeleteUser(id int) (bool, error) {
	session := mongo.GetMongoSession()
	defer session.Close()
	userCollection := session.DB(database).C(collection)
	err := userCollection.Remove(bson.M{"id": id})
	if err != nil {
		return false, err
	}
	return true, nil
}
