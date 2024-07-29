package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: creating user"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Getting all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Getting user by id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Updating user informations"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Deleting user"))
}
