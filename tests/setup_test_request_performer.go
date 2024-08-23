package tests

import (
	"api/internal/models"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func PerformAuthenticatedRequest(method string, url string, body []byte) (*httptest.ResponseRecorder, error) {
	databaseConnection := GetTestDBConnection()

	router := GetTestRouter()

	databaseConnection.AutoMigrate(&models.User{})

	bearerToken, err := PerformLogin()

	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	return requestRecorder, nil
}

func PerformUnAuthenticatedRequest(method string, url string, body []byte) (*httptest.ResponseRecorder, error) {
	GetTestDBConnection()

	router := GetTestRouter()

	request, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	return requestRecorder, nil
}
