package tests

import (
	"api/src/models"
	"api/src/services"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"gorm.io/gorm"
)

var username string = "testuser"
var email string = "test@example.com"
var password string = "password"
var testUser *models.User

func CreateTestUser(db *gorm.DB) *models.User {
	if testUser != nil && testUser.ID != 0 {
		return testUser
	}

	request := db.Where("email = ?", email)

	if request.RowsAffected > 0 {
		db.Where("email = ?", email).Take(&testUser)

		return testUser
	}

	hashedPassword, _ := services.HashPassword(password)
	testUser = &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err := db.Create(testUser).Error

	if err != nil {
		panic(err)
	}

	return testUser
}

func PerformLogin() (string, error) {
	if testUser == nil {
		CreateTestUser(services.GetConnection())
	}

	loginCredentials := map[string]string{
		"email":    testUser.Email,
		"password": password,
	}

	body, _ := json.Marshal(loginCredentials)

	request, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	if requestRecorder.Code != http.StatusOK {
		return "", fmt.Errorf("PerformLogin(): failed with status: %d", requestRecorder.Code)
	}

	var response map[string]string

	json.Unmarshal(requestRecorder.Body.Bytes(), &response)

	token, ok := response["token"]

	if !ok {
		return "", fmt.Errorf("PerformLogin(): no token found in response")
	}

	return token, nil
}
