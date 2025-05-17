package main

import (
	"day1-go-mvc/controllers"
	"net/http"
)

func main() {
	userController := &controllers.UserController{}
	http.HandleFunc("/users", userController.ListUsers)
	port := ":8080"
	println("服务器运行在 http://localhost" + port)
	http.ListenAndServe(port, nil)
}
