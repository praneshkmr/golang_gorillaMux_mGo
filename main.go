package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"./app/dao/mongo"
	UserRoutes "./app/routes/user"
	Config "./app/util/config"
)

func main() {
	config := Config.LoadConfig("./config/local.json")
	mongo.SetConfig(config)

	router := mux.NewRouter().StrictSlash(true)
	UserRoutes.AssignRoutes(router)

	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
