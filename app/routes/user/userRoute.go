package user

import (
	"encoding/json"
	"net/http"

	"./../../../app/model"
	"github.com/gorilla/mux"

	UserService "./../../../app/service/user"
)

func addUser(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	addedUser, err := UserService.AddUser(&user)
	if err != nil {
		panic(err)
	}
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(addedUser); err != nil {
		panic(err)
	}
}

// AssignRoutes is used to Setup the REST routes with the appropriate Handlers
func AssignRoutes(router *mux.Router) {
	router.HandleFunc("/user", addUser).Methods("POST")
	http.ListenAndServe(":3000", router)
}
