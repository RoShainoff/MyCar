package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/service/mock"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCarHandler_CreateCar(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.CarServiceMock{
		CreateCarFunc: func(newCar *car.Car, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	handler := NewCarHandler(mockService)

	router := gin.New()
	router.POST("/api/car", func(c *gin.Context) {
		c.Set("UserId", uuid.New())
		handler.CreateCar(c)
	})

	reqBody := CarRq{
		Brand:            car.AlfaRomeo,
		DriveType:        car.FWD,
		BodyType:         car.Sedan,
		TransmissionType: car.TransmissionTypeAutomatic,
		FuelType:         vehicle.Petrol,
		Year:             2020,
		Plate:            "A123BC",
		Vin:              "VIN1234567890",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/car", bytes.NewBuffer(body))
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

func TestCarHandler_GetCarById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockCar := car.NewCar(mockedId, uuid.New(), vehicle.Petrol, car.AlfaRomeo, 2020, "A123BC", car.FWD, car.Sedan, car.TransmissionTypeAutomatic)
	mockService := &mock.CarServiceMock{
		GetCarByIdFunc: func(id, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
			return mockCar, nil
		},
	}
	handler := NewCarHandler(mockService)
	router := gin.New()
	router.GET("/api/car/:id", func(c *gin.Context) {
		c.Set("UserId", uuid.New())
		handler.GetCarById(c)
	})

	req, _ := http.NewRequest("GET", "/api/car/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
	var resp car.Car
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp.GetId() != mockedId {
		t.Error("id не совпадает")
	}
}

func TestCarHandler_UpdateCar(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	mockService := &mock.CarServiceMock{
		UpdateCarFunc: func(id, userId uuid.UUID, updatedCar *car.Car) *model.ApplicationError {
			return nil
		},
	}
	handler := NewCarHandler(mockService)
	router := gin.New()
	router.PUT("/api/car/:id", func(c *gin.Context) {
		c.Set("UserId", uuid.New())
		handler.UpdateCar(c)
	})

	reqBody := CarRq{
		Brand:            car.AlfaRomeo,
		DriveType:        car.FWD,
		BodyType:         car.Sedan,
		TransmissionType: car.TransmissionTypeAutomatic,
		FuelType:         vehicle.Petrol,
		Year:             2020,
		Plate:            "A123BC",
		Vin:              "VIN1234567890",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/api/car/"+mockedId.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}

func TestCarHandler_DeleteCar(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	userId := uuid.New()
	mockService := &mock.CarServiceMock{
		DeleteCarFunc: func(id, uId uuid.UUID) *model.ApplicationError {
			return nil
		},
	}
	handler := NewCarHandler(mockService)
	router := gin.New()
	router.DELETE("/api/car/:id", func(c *gin.Context) {
		c.Set("UserId", userId)
		handler.DeleteCar(c)
	})

	req, _ := http.NewRequest("DELETE", "/api/car/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}
