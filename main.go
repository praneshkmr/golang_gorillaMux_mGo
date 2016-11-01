package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	UserRoutes "./app/routes/user"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	UserRoutes.AssignRoutes(router)
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
