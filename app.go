package main

import "github.com/jutionck/golang-todo-apps/delivery"

// @title           Todo App
// @version         1.0

// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @schemes http
func main() {
	delivery.NewServer().Run()
}
