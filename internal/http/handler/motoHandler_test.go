package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/moto"
	"MyCar/internal/service/mock"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMotoHandler_CreateMoto(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	userId := uuid.New()
	mockService := &mock.MotoServiceMock{
		CreateMotoFunc: func(newMoto *moto.Moto, uId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
			if uId != userId {
				t.Fatalf("userId не совпадает")
			}
			return mockedId, nil
		},
	}
	handler := NewMotoHandler(mockService)
	router := gin.New()
	router.POST("/api/moto", func(c *gin.Context) {
		c.Set("UserId", userId)
		handler.CreateMoto(c)
	})

	reqBody := MotoRq{
		Brand:            moto.Minsk,
		Category:         moto.Classic,
		TransmissionType: moto.TransmissionTypeManual,
		FuelType:         vehicle.Petrol,
		Year:             2022,
		Plate:            "MOTO123",
		Vin:              "VINMOTO123456",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/moto", bytes.NewBuffer(body))
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
	if resp["id"] == nil {
		t.Fatalf("id не найден в ответе")
	}
}

func TestMotoHandler_GetMotoById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	userId := uuid.New()
	mockMoto := moto.NewMoto(mockedId, userId, vehicle.Petrol, moto.Minsk, 2022, "MOTO123", moto.Classic, moto.TransmissionTypeManual)
	mockService := &mock.MotoServiceMock{
		GetMotoByIdFunc: func(id, uId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
			if id != mockedId || uId != userId {
				t.Fatalf("id или userId не совпадают")
			}
			return mockMoto, nil
		},
	}
	handler := NewMotoHandler(mockService)
	router := gin.New()
	router.GET("/api/moto/:id", func(c *gin.Context) {
		c.Set("UserId", userId)
		handler.GetMotoById(c)
	})

	req, _ := http.NewRequest("GET", "/api/moto/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
	var resp moto.Moto
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("ошибка разбора ответа: %v", err)
	}
	if resp.GetId() != mockedId {
		t.Fatalf("id не совпадает")
	}
}

func TestMotoHandler_UpdateMoto(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	userId := uuid.New()
	mockService := &mock.MotoServiceMock{
		UpdateMotoFunc: func(id, uId uuid.UUID, updatedMoto *moto.Moto) *model.ApplicationError {
			if id != mockedId || uId != userId {
				t.Fatalf("id или userId не совпадают")
			}
			return nil
		},
	}
	handler := NewMotoHandler(mockService)
	router := gin.New()
	router.PUT("/api/moto/:id", func(c *gin.Context) {
		c.Set("UserId", userId)
		handler.UpdateMoto(c)
	})

	reqBody := MotoRq{
		Brand:            moto.Minsk,
		Category:         moto.Classic,
		TransmissionType: moto.TransmissionTypeManual,
		FuelType:         vehicle.Petrol,
		Year:             2022,
		Plate:            "MOTO123",
		Vin:              "VINMOTO123456",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/api/moto/"+mockedId.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}

func TestMotoHandler_DeleteMoto(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedId := uuid.New()
	userId := uuid.New()
	mockService := &mock.MotoServiceMock{
		DeleteMotoFunc: func(id, uId uuid.UUID) *model.ApplicationError {
			if id != mockedId || uId != userId {
				t.Fatalf("id или userId не совпадают")
			}
			return nil
		},
	}
	handler := NewMotoHandler(mockService)
	router := gin.New()
	router.DELETE("/api/moto/:id", func(c *gin.Context) {
		c.Set("UserId", userId)
		handler.DeleteMoto(c)
	})

	req, _ := http.NewRequest("DELETE", "/api/moto/"+mockedId.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался статус 200, получен %d", w.Code)
	}
}
