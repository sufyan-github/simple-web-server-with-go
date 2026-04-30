package services

import (
	"go-web-server/models"
	"go-web-server/repository"
)

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func CreateUser(user models.User) error {
	return repository.CreateUser(user)
}

func GetUserByID(id int) (models.User, error) {
	return repository.GetUserByID(id)
}

func UpdateUser(user models.User) error {
	return repository.UpdateUser(user)
}

func DeleteUser(id int) error {
	return repository.DeleteUser(id)
}	

