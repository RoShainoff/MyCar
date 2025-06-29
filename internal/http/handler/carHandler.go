package handler

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CarHandler struct {
	service service.AbstractCarService
}

type CarRq struct {
	Brand            car.Brand                `json:"brand" binding:"required"`
	DriveType        car.DriveTypeKind        `json:"drive_type" binding:"required"`
	BodyType         car.BodyTypeKind         `json:"body_type" binding:"required"`
	TransmissionType car.TransmissionTypeKind `json:"transmission_type" binding:"required"`
	FuelType         vehicle.FuelType         `json:"fuel_type" binding:"required"`
	Year             int                      `json:"year" binding:"required"`
	Plate            string                   `json:"plate" binding:"required"`
	Vin              string                   `json:"vin" binding:"required"`
}

func NewCarHandler(s service.AbstractCarService) CarHandler {
	return CarHandler{service: s}
}

// CreateCar godoc
// @Summary Create a new car
// @Description Creates a new car for the authenticated user
// @Tags cars
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body CarRq true "Car creation data"
// @Success 200 {object} map[string]interface{} "Returns ID of created car"
// @Failure 400 {object} response "Invalid request data"
// @Failure 401 {object} response "Unauthorized"
// @Failure 500 {object} response "Internal server error"
// @Router /api/car [post]
func (h *CarHandler) CreateCar(c *gin.Context) {
	var req CarRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	newCar := car.NewCar(uuid.Nil, userId, req.FuelType, req.Brand, req.Year, req.Plate, req.DriveType, req.BodyType, req.TransmissionType)

	if err := newCar.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return
	}

	id, err := h.service.CreateCar(newCar, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetCarById godoc
// @Summary Получить автомобиль по ID
// @Description Возвращает автомобиль по идентификатору для аутентифицированного пользователя
// @Tags cars
// @Produce json
// @Security BearerAuth
// @Param id path string true "Car ID"
// @Success 200 {object} car.Car "Данные автомобиля"
// @Failure 400 {object} response "Некорректный ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Автомобиль не найден"
// @Failure 500 {object} response "Внутренняя ошибка сервера"
// @Router /api/car/{id} [get]
func (h *CarHandler) GetCarById(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Некорректный car ID")
		return
	}
	carObj, getErr := h.service.GetCarById(id, userId)
	if getErr != nil {
		errorResponse(c, http.StatusNotFound, getErr.Error())
		return
	}
	c.JSON(http.StatusOK, carObj)
}

// UpdateCar godoc
// @Summary Update a car
// @Description Updates an existing car for the authenticated user
// @Tags cars
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Car ID"
// @Param input body CarRq true "Car update data"
// @Success 200 "Car updated successfully"
// @Failure 400 {object} response "Invalid request data or ID"
// @Failure 401 {object} response "Unauthorized"
// @Failure 404 {object} response "Car not found"
// @Failure 500 {object} response "Internal server error"
// @Router /api/car/{id} [put]
func (h *CarHandler) UpdateCar(c *gin.Context) {
	var req CarRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid car ID")
		return
	}

	updatedCar := car.NewCar(id, userId, req.FuelType, req.Brand, req.Year, req.Plate, req.DriveType, req.BodyType, req.TransmissionType)

	if err := updatedCar.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return
	}

	updateErr := h.service.UpdateCar(id, userId, updatedCar)
	if updateErr != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteCar godoc
// @Summary Delete a car
// @Description Deletes an existing car for the authenticated user
// @Tags cars
// @Produce json
// @Security BearerAuth
// @Param id path string true "Car ID"
// @Success 200 "Car deleted successfully"
// @Failure 400 {object} response "Invalid car ID"
// @Failure 401 {object} response "Unauthorized"
// @Failure 404 {object} response "Car not found"
// @Failure 500 {object} response "Internal server error"
// @Router /api/car/{id} [delete]
func (h *CarHandler) DeleteCar(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid car ID")
		return
	}
	err = h.service.DeleteCar(id, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
