package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	_ "github.com/lib/pq"
	"io"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/hash"
	"nandes007/blog-post-rest-api/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

func TestLoginSuccess(t *testing.T) {
	db := test.SetupTestDB()
	truncateUser(db)
	router := test.SetupRouter(db)
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
	db := test.SetupTestDB()
	truncateUser(db)
	router := test.SetupRouter(db)
	createUser(db)
	token := test.GenerateToken()

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
	db := test.SetupTestDB()
	truncateUser(db)
	router := test.SetupRouter(db)
	createUser(db)
	token := test.GenerateToken()

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
