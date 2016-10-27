package main

import (
	"log"

	"./app/dao/mongo"
	"./app/model"
)

func main() {
	user := model.User{ID: 1, Name: "Pranesh Kumar", Email: "praneshkmr@gmail.com", Password: "temp123"}
	_, err := mongo.AddUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted")

	user2, err := mongo.GetUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user2)

	user.Password = "temp1234"
	user3, err := mongo.UpdateUser(user.ID, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user3)

	isDeleted, err := mongo.DeleteUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(isDeleted)
}
