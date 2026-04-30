package main

import (
	"net/http"
	"go-web-server/config"
	"go-web-server/routes"
	_ "go-web-server/docs"
)


// @title Go Web Server API
// @version 1.0
// @description This is a sample server in Go
// @host localhost:8080
// @BasePath /

func main() {
	config.ConnectDB()
	routes.RegisterRoutes()

	http.ListenAndServe(":8080", nil)
}