package handlers

import (
	"encoding/json"
	"net/http"
	"go-web-server/models"
	"strconv"
	"fmt"
)

var users = []models.User{
	{ID: 1, Name: "Alice", Email: "alice@gmail.com"},
	{ID: 2, Name: "Bob", Email: "bob@gmail.com"},
}

// Helper function to write JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// home routes
func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// Handle /users route(GET, POST)
func UsersHandler(w http.ResponseWriter, r *http.Request){

	switch r.Method {

		case http.MethodGet:
			respondJSON(w, http.StatusOK, users)

		case http.MethodPost:
			var newUser models.User

			err := json.NewDecoder(r.Body).Decode(&newUser)
			if err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}

			users = append(users, newUser)
			respondJSON(w, http.StatusCreated, newUser)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handle /user?id=1 route(GET,PUT,DELETE)
func UserHandler(w http.ResponseWriter, r *http.Request){

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	switch r.Method {

	case http.MethodGet:
		for _, u := range users {
			if u.ID == id {
				respondJSON(w, http.StatusOK, u)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)


	case http.MethodPut:

		var updatedUser models.User

		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		for i, u := range users {
			if u.ID == id {
				updatedUser.ID = id
				users[i] = updatedUser
				respondJSON(w, http.StatusOK, updatedUser)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)


	case http.MethodDelete:
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}