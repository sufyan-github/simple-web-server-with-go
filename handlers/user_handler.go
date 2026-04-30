package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"go-web-server/models"
	"go-web-server/services"
)


// GetUsers godoc
// @Summary Get all users
// @Description Retrieve all users from database
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := services.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}


// CreateUser godoc
// @Summary Create new user
// @Description Add a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Router /user/create [post]

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	services.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}


// GetUser godoc
// @Summary Get single user
// @Description Get user by ID
// @Tags Users
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} models.User
// @Router /user/get [get]

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	user, _ := services.GetUserByID(id)
	json.NewEncoder(w).Encode(user)
}


// UpdateUser godoc
// @Summary Update user
// @Description Update user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Param user body models.User true "Updated Data"
// @Success 200 {object} models.User
// @Router /user/update [put]

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = id

	services.UpdateUser(user)
	json.NewEncoder(w).Encode(user)
}


// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags Users
// @Param id query int true "User ID"
// @Success 204
// @Router /user/delete [delete]

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	services.DeleteUser(id)
	w.WriteHeader(http.StatusNoContent)
}