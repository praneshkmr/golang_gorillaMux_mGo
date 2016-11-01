package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"./../../../app/model"
	"github.com/gorilla/mux"

	UserService "./../../../app/service/user"
)

func addUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Failed to Decode Request Body"))
		return
	}
	addedUser, err := UserService.AddUser(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Failed to Add User"))
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(addedUser); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Failed to Encode Response Body"))
	}
}

func getUserByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Usage -> /user/:ID where ID is an Integer"))
		return
	}
	user, err := UserService.GetUserByID(userID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(user); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Failed to Encode Response Body"))
	}
}

func updateUserByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Usage -> /user/:ID where ID is an Integer"))
		return
	}
	decoder := json.NewDecoder(req.Body)
	var user model.User
	err = decoder.Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Failed to Decode Request Body"))
		return
	}
	updatedUser, err := UserService.UpdateUserByID(userID, &user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Failed to Update User"))
		return
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(updatedUser); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Failed to Encode Response Body"))
	}
}

func deleteUserByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Usage -> /user/:ID where ID is an Integer"))
		return
	}
	_, err = UserService.DeleteUserByID(userID)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Failed to Delete User"))
		return
	}
	res.WriteHeader(http.StatusOK)
}

// AssignRoutes is used to Setup the REST routes with the appropriate Handlers
func AssignRoutes(router *mux.Router) {
	router.HandleFunc("/user", addUser).Methods("POST")
	router.HandleFunc("/user/{id}", getUserByID).Methods("GET")
	router.HandleFunc("/user/{id}", updateUserByID).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUserByID).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}
