package main

import (
	"log"

	Counter "./app/dao/mongo/counter"
	User "./app/dao/mongo/user"
	"./app/model"
)

func main() {
	id, err := Counter.GetNextCounter("user")
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{ID: id, Name: "Pranesh Kumar", Email: "praneshkmr@gmail.com", Password: "temp123"}
	_, err = User.AddUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted")

	user2, err := User.GetUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user2)

	user.Password = "temp1234"
	user3, err := User.UpdateUser(user.ID, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user3)

	isDeleted, err := User.DeleteUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(isDeleted)
}
