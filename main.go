package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"nandes007/blog-post-rest-api/app"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/env"
	"nandes007/blog-post-rest-api/middleware"
	"nandes007/blog-post-rest-api/repository"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"time"
)

func main() {
	env.Load()
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository, userRepository, db, validate)
	postController := controller.NewPostController(postService)
	router := app.NewRouter(userController, postController)

	server := http.Server{
		Addr:        ":9001",
		Handler:     middleware.NewHandler(router),
		ReadTimeout: 5 * time.Second,
	}

	fmt.Printf("Server is running at http://localhost%s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
