package controllers

import "net/http"

// CreateUser create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
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
