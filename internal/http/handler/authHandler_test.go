package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/service/mock"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &mock.AuthServiceMock{
		AuthUserFunc: func(login, password string) (string, *model.ApplicationError) {
			return "mocked-jwt-token", nil
		},
	}
	handler := NewAuthHandler(mockService)

	router := gin.New()
	router.POST("/api/auth/login", handler.Login)

	reqBody := AuthReq{
		Login:    "user123",
		Password: "securePassword123",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp["token"] != "mocked-jwt-token" {
		t.Error("токен не совпадает")
	}
}

func TestAuthHandler_Login_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &mock.AuthServiceMock{}
	handler := NewAuthHandler(mockService)

	router := gin.New()
	router.POST("/api/auth/login", handler.Login)

	// Не хватает поля Password
	reqBody := map[string]string{"Login": "user123"}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("ожидался статус 400, получен %d", w.Code)
	}
}
