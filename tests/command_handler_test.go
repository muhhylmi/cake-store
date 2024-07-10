package tests

import (
	"cake-store/modules/cakes"
	"cake-store/modules/cakes/handlers"
	"cake-store/modules/cakes/repositories"
	"cake-store/modules/cakes/usecases"
	"cake-store/utils/config"
	"cake-store/utils/logger"
	"cake-store/utils/middleware"
	"cake-store/utils/wrapper"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"gopkg.in/go-playground/assert.v1"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB(config *config.Configurations) *sql.DB {
	db, err := sql.Open(config.DB_DIALECT, config.DB_URI)
	wrapper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB, config *config.Configurations) http.Handler {
	validate := validator.New()
	logger := logger.Newlogger()

	repository := repositories.NewRepository(logger, db)
	usecase := usecases.NewUsecase(logger, repository)
	categoryController := handlers.NewCakeHandler(logger, usecase, validate)

	router := cakes.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router, config)
}

func truncateCake(db *sql.DB) {
	db.Exec("TRUNCATE cakes")
}

func createCake(db *sql.DB) {
	db.Exec("insert into cakes(id, title,description,rating,image) values (100, 'test cake', 'mantap', 8.5, 'image.jpg')")
}

func TestCakeSuccess(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)

	router := setupRouter(db, config)
	requestBody := strings.NewReader(`{"title": "test cake","description":"test","image":"test.jpg","rating":8}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test cake", responseBody["data"].(map[string]interface{})["title"])
}

func TestCreateCakeFailed(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)

	router := setupRouter(db, config)
	requestBody := strings.NewReader(`{"title": "","description":"test","image":"test.jpg","rating":8}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST ERROR", responseBody["status"])
}

func TestGetCakeByIdSuccess(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)
	createCake(db)

	router := setupRouter(db, config)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/cakes/100", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test cake", responseBody["data"].(map[string]interface{})["title"])
}

func TestGetCakeByIdFailed(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)
	createCake(db)

	router := setupRouter(db, config)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/cakes/99", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND ERROR", responseBody["status"])
}

func TestListCakeSuccess(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)
	createCake(db)

	router := setupRouter(db, config)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/cakes", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	dataArray := responseBody["data"].([]interface{})
	firstDataElement := dataArray[0].(map[string]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test cake", firstDataElement["title"])
}

func TestListCakeEmpty(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)

	router := setupRouter(db, config)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/cakes", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	dataArray := responseBody["data"].([]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, 0, len(dataArray))
}

func TestDeleteCakeByIdSuccess(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)
	createCake(db)

	router := setupRouter(db, config)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/cakes/100", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, true, responseBody["data"].(bool))
}

func TestDeleteCakeByIdFailed(t *testing.T) {
	config := config.GetConfig()
	db := setupTestDB(config)
	truncateCake(db)
	createCake(db)

	router := setupRouter(db, config)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/cakes/99", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("API-Key", config.API_KEY)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND ERROR", responseBody["status"])
}
