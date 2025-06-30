package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/service/mock"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.UserServiceMock{
		CreateUserFunc: func(name, surname, email, login, password string) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	handler := NewUserHandler(mockService)
	router := gin.New()
	router.POST("/api/user", func(c *gin.Context) {
		handler.CreateUser(c)
	})

	reqBody := map[string]string{
		"name":     "Ivan",
		"surname":  "Ivanov",
		"email":    "ivan@example.com",
		"login":    "ivan1234",
		"password": "password1!",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("ожидался статус 201, получен %d, тело: %s", w.Code, w.Body.String())
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp["id"] == nil {
		t.Error("id не возвращён")
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockUser := &auth.User{
		Id:      mockedId,
		Login:   "user123",
		Name:    "John",
		Surname: "Doe",
	}
	mockService := &mock.UserServiceMock{
		GetUserFunc: func(id uuid.UUID) (*auth.User, *model.ApplicationError) {
			return mockUser, nil
		},
	}
	handler := NewUserHandler(mockService)
	router := gin.New()
	router.GET("/api/user", func(c *gin.Context) {
		c.Set("UserId", mockedId)
		handler.GetUser(c)
	})

	req, _ := http.NewRequest("GET", "/api/user", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
	var resp struct {
		User UserRsp `json:"user"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp.User.Id != mockedId {
		t.Errorf("ожидался id %v, получен %v", mockedId, resp.User.Id)
	}
	if resp.User.Login != mockUser.Login {
		t.Errorf("ожидался login %v, получен %v", mockUser.Login, resp.User.Login)
	}
	if resp.User.Name != mockUser.Name {
		t.Errorf("ожидалось имя %v, получено %v", mockUser.Name, resp.User.Name)
	}
	if resp.User.Surname != mockUser.Surname {
		t.Errorf("ожидалась фамилия %v, получена %v", mockUser.Surname, resp.User.Surname)
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.UserServiceMock{
		UpdateUserFunc: func(id uuid.UUID, name, surname, email, login, password string) *model.ApplicationError {
			return nil
		},
	}
	handler := NewUserHandler(mockService)
	router := gin.New()
	router.PUT("/api/user/:id", func(c *gin.Context) {
		c.Set("UserId", mockedId)
		handler.UpdateUser(c)
	})

	reqBody := map[string]string{
		"name":     "Ivan",
		"surname":  "Ivanov",
		"email":    "ivan@example.com",
		"login":    "ivan123",
		"password": "pass",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/api/user/"+mockedId.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.UserServiceMock{
		DeleteUserFunc: func(id uuid.UUID) *model.ApplicationError {
			return nil
		},
	}
	handler := NewUserHandler(mockService)
	router := gin.New()
	router.DELETE("/api/user/:id", func(c *gin.Context) {
		c.Set("UserId", mockedId)
		handler.DeleteUser(c)
	})

	req, _ := http.NewRequest("DELETE", "/api/user/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}
