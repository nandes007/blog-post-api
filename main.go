package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"nandes007/blog-post-rest-api/app"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/middleware"
	"nandes007/blog-post-rest-api/repository"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"time"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	server := http.Server{
		Addr:        ":9001",
		Handler:     middleware.NewHandler(router),
		ReadTimeout: 5 * time.Second,
	}

	fmt.Printf("Server is running at http://localhost%s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
