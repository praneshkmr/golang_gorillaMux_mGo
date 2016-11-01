# Go Lang REST Server using Gorilla Mux and mGo

A REST Server created using Go Lang with the Gorilla Mux as the Router and MongoDB as the Database.
MGo driver is used to connect to mongoDB from the Server

### Tech

This REST Server uses the Following open source projects to work properly:

* [Gorilla Mux] - A powerful URL router and dispatcher for golang
* [mGo] - The MongoDB driver for Go
* [Negroni] - Idiomatic HTTP Middleware for Golang

### Start

Provide the MongoDB Host details in the Configuration File located under **conf** folder

To start the server, install the Dependencies and run the following command.

```sh
$ go run main.go
```

### Test

To test the User CRUD REST API's, run the following command.

```sh
$ go test test/user_api_test.go 
```

   [Gorilla Mux]: <https://github.com/gorilla/mux>
   [mGo]: <https://github.com/go-mgo/mgo>
   [Negroni]: <https://github.com/urfave/negroni>
   
