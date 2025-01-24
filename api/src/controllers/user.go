package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewUserRepository(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	responses.JSON(w, http.StatusCreated, user)
	// w.Write([]byte(fmt.Sprintf("Inserted id: %d", userId)))
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
