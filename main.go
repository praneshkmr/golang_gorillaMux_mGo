package main

import (
	"log"

	"./app/model"
	UserService "./app/service/user"
)

func main() {
	user := model.User{Name: "Pranesh Kumar", Email: "praneshkmr@gmail.com", Password: "temp123"}
	_, err := UserService.AddUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted")

	user2, err := UserService.GetUserByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user2)

	user.Password = "temp1234"
	user3, err := UserService.UpdateUserByID(user.ID, &user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user3)

	isDeleted, err := UserService.DeleteUserByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(isDeleted)
}
