package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/middleware"
	"nandes007/blog-post-rest-api/repository"
	"nandes007/blog-post-rest-api/routes"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func SetupTestDB() *sql.DB {
	connStr := "postgres://postgres:postgre@localhost/blog_post_test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func SetupRouter(db *sql.DB) http.Handler {
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

	return middleware.NewHandler(router)
}

func GenerateToken() string {
	db := SetupTestDB()
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"email": "test@example.com", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9001/api/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	token := responseBody["data"]
	return token.(string)
}
