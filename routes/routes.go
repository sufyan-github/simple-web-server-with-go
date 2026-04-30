package routes

import(
	"net/http"
	"go-web-server/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("/",handlers.HomeHandler)
	http.HandleFunc("/users",handlers.UsersHandler)
	http.HandleFunc("/user",handlers.UserHandler)

	// Static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

}