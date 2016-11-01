package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"strconv"

	UserRoute "./../app/routes/user"
)

var (
	server   *httptest.Server
	reader   io.Reader
	usersURL string
	id       string
)

type CreateUserResponse struct {
	ID int `json:"id"`
}

func init() {
	r := mux.NewRouter()
	UserRoute.AssignRoutes(r)
	server = httptest.NewServer(r) //Creating new server with the user handlers

	usersURL = fmt.Sprintf("%s/users/", server.URL) //Grab the address for the API endpoint
}

func TestCreateUser(t *testing.T) {
	userJSON := `{"name": "Pranesh", "email": "praneshkmr@gmail.com", "password": "temp123"}`

	reader = strings.NewReader(userJSON)                      //Convert string to reader
	request, err := http.NewRequest("POST", usersURL, reader) //Create request with JSON body
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 201 {
		t.Errorf("Created expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}

	var responseJSON CreateUserResponse
	json.NewDecoder(res.Body).Decode(&responseJSON)
	id = strconv.Itoa(responseJSON.ID)
}

func TestGetUserByID(t *testing.T) {
	request, err := http.NewRequest("GET", usersURL+id, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUpdateUserByID(t *testing.T) {
	updatedUserJSON := `{"name": "Pranesh", "email": "praneshkmr@gmail.com", "password": "temp1234"}`

	reader = strings.NewReader(updatedUserJSON)
	request, err := http.NewRequest("PUT", usersURL+id, reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestDeleteUserByID(t *testing.T) {
	request, err := http.NewRequest("DELETE", usersURL+id, nil)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
