package routes

import (
	"net/http"
	"go-web-server/handlers"
	_ "go-web-server/docs"   
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes() {
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/user/create", handlers.CreateUser)
	http.HandleFunc("/user/get", handlers.GetUser)
	http.HandleFunc("/user/update", handlers.UpdateUser)
	http.HandleFunc("/user/delete", handlers.DeleteUser)
	http.Handle("/swagger/", httpSwagger.WrapHandler)	

}