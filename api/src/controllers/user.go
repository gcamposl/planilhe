package controllers

import (
	"api/src/database"
	"api/src/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	buildUser(body)

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

}

// GetAllUsers returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

// GetUser return specific user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

// UpdateUser update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

// DeleteUser delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

func buildUser(body []byte) models.User {
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	return user
}
