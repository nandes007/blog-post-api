package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"io"
	"nandes007/blog-post-rest-api/app"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/hash"
	"nandes007/blog-post-rest-api/middleware"
	"nandes007/blog-post-rest-api/repository"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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

// Must be move for initialize application!.
func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository, userRepository, db, validate)
	postController := controller.NewPostController(postService)
	router := app.NewRouter(userController, postController)

	return middleware.NewHandler(router)
}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE TABLE users")
}

func createUser(db *sql.DB) {
	currentDate := helper.GetCurrentTime()
	passwordHashed := hash.PasswordHash("password")
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, "test", "test@example.com", passwordHashed, currentDate, currentDate)

	if err != nil {
		fmt.Println("Error")
	}

	return
}

func generateToken() string {
	db := SetupTestDB()
	router := setupRouter(db)

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

func TestLoginSuccess(t *testing.T) {
	db := SetupTestDB()
	truncateUser(db)
	router := setupRouter(db)
	createUser(db)

	requestBody := strings.NewReader(`{"email": "test@example.com", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9001/api/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
}

func TestCreateUserSuccess(t *testing.T) {
	db := SetupTestDB()
	truncateUser(db)
	router := setupRouter(db)
	createUser(db)
	token := generateToken()

	requestBody := strings.NewReader(`{"name": "test2", "email": "test2@example.com", "password": "password"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9001/api/users", requestBody)
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "test2", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "test2@example.com", responseBody["data"].(map[string]interface{})["email"])
}

func TestGetAllUserSuccess(t *testing.T) {
	db := SetupTestDB()
	truncateUser(db)
	router := setupRouter(db)
	createUser(db)
	token := generateToken()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:9001/api/users", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	users := responseBody["data"].([]interface{})
	assert.Equal(t, "test", users[0].(map[string]interface{})["name"])
	assert.Equal(t, "test@example.com", users[0].(map[string]interface{})["email"])
}
