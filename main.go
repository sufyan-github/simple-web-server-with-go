package main

import (
	"fmt"
	"net/http"
	"go-web-server/routes"
)

func main(){
	fmt.Println("Starting the server on port 8080...")

	// Load routes
	routes.RegisterRoutes()

	// Start the server
	http.ListenAndServe(":8080",nil)
}