package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle"
	"MyCar/internal/service/mock"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestExpenseHandler_CreateExpense(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.ExpenseServiceMock{
		CreateExpenseFunc: func(newExpense *expense.Expense, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	handler := NewExpenseHandler(mockService)
	router := gin.New()
	router.POST("/api/expense", func(c *gin.Context) {
		c.Set("UserId", uuid.New())
		handler.CreateExpense(c)
	})

	reqBody := ExpenseRq{
		VehicleId:    uuid.New(),
		VehicleType:  vehicle.Car,
		Category:     expense.Fuel,
		Amount:       100.5,
		Currency:     "RUB",
		ExchangeRate: 1.0,
		Date:         time.Now().Format("2006-01-02"),
		Note:         "Заправка",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/expense", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d, тело: %s", w.Code, w.Body.String())
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp["id"] == nil {
		t.Error("id не возвращён")
	}
}

func TestExpenseHandler_GetExpenseById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockExpense := expense.NewExpense(mockedId, uuid.New(), vehicle.Car, expense.Fuel, 100.5, "RUB", 1.0, time.Now(), "Заправка")
	mockService := &mock.ExpenseServiceMock{
		GetExpenseByIdFunc: func(id, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
			return mockExpense, nil
		},
	}
	handler := NewExpenseHandler(mockService)
	router := gin.New()
	router.GET("/api/expense/:id", func(c *gin.Context) {
		c.Set("UserId", uuid.New())
		handler.GetExpenseById(c)
	})

	req, _ := http.NewRequest("GET", "/api/expense/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
	var resp expense.Expense
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp.GetId() != mockedId {
		t.Error("id не совпадает")
	}
}

func TestExpenseHandler_UpdateExpense(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.ExpenseServiceMock{
		UpdateExpenseFunc: func(id, userId uuid.UUID, updatedExpense *expense.Expense) *model.ApplicationError {
			return nil
		},
	}
	handler := NewExpenseHandler(mockService)
	router := gin.New()
	router.PUT("/api/expense/:id", func(c *gin.Context) {
		c.Set("UserId", uuid.New())
		handler.UpdateExpense(c)
	})

	reqBody := ExpenseRq{
		VehicleId:    uuid.New(),
		VehicleType:  vehicle.Car,
		Category:     expense.Fuel,
		Amount:       200.0,
		Currency:     "RUB",
		ExchangeRate: 1.0,
		Date:         time.Now().Format("2006-01-02"),
		Note:         "Обновление",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/api/expense/"+mockedId.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}

func TestExpenseHandler_DeleteExpense(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	userId := uuid.New()
	mockService := &mock.ExpenseServiceMock{
		DeleteExpenseFunc: func(id, uId uuid.UUID) *model.ApplicationError {
			return nil
		},
	}
	handler := NewExpenseHandler(mockService)
	router := gin.New()
	router.DELETE("/api/expense/:id", func(c *gin.Context) {
		c.Set("UserId", userId)
		handler.DeleteExpense(c)
	})

	req, _ := http.NewRequest("DELETE", "/api/expense/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}
