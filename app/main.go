package main

import (
	"fmt"
	app "nandes007/blog-post-rest-api/config"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/middleware"
	"nandes007/blog-post-rest-api/repository"
	"nandes007/blog-post-rest-api/routes"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	// env.Load()
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository, userRepository, db, validate)
	postController := controller.NewPostController(postService)

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db, validate)
	authController := controller.NewAuthController(authService)
	router := routes.NewRouter(userController, postController, authController)

	server := http.Server{
		Addr:        ":9001",
		Handler:     middleware.NewHandler(router),
		ReadTimeout: 5 * time.Second,
	}

	fmt.Printf("Server is running at http://localhost%s\n", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
