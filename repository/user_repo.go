package repository

import (
	"go-web-server/config"
	"go-web-server/models"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	return users, nil
}

// CreateUser adds a new user to the database
func CreateUser(user models.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(id int) (models.User, error) {
	var u models.User

	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)

	return u, err
}

// UpdateUser updates an existing user in the database
func UpdateUser(user models.User) error {
	_, err := config.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	return err
}

// DeleteUser removes a user from the database
func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}