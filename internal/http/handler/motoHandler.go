package handler

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/moto"
	"MyCar/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type MotoHandler struct {
	service service.AbstractMotoService
}

type MotoRq struct {
	Brand            moto.BrandKind            `json:"brand" binding:"required"`
	Category         moto.CategoryKind         `json:"category" binding:"required"`
	TransmissionType moto.TransmissionTypeKind `json:"transmission_type" binding:"required"`
	FuelType         vehicle.FuelType          `json:"fuel_type" binding:"required"`
	Year             int                       `json:"year" binding:"required"`
	Plate            string                    `json:"plate" binding:"required"`
	Vin              string                    `json:"vin" binding:"required"`
}

func NewMotoHandler(s service.AbstractMotoService) MotoHandler {
	return MotoHandler{service: s}
}

// CreateMoto godoc
// @Summary Create a new moto
// @Description Creates a new moto for the authenticated user
// @Tags motos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body MotoRq true "Moto creation data"
// @Success 200 {object} map[string]interface{} "Returns ID of created moto"
// @Failure 400 {object} response "Invalid request data"
// @Failure 401 {object} response "Unauthorized"
// @Failure 500 {object} response "Internal server error"
// @Router /api/moto [post]
func (h *MotoHandler) CreateMoto(c *gin.Context) {
	var req MotoRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	newMoto := moto.NewMoto(uuid.Nil, userId, req.FuelType, req.Brand, req.Year, req.Plate, req.Category, moto.TransmissionTypeManual)
	newMoto.SetVin(req.Vin)

	if err := newMoto.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return
	}

	id, err := h.service.CreateMoto(newMoto, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetMotoById godoc
// @Summary Получить мотоцикл по ID
// @Description Возвращает мотоцикл по идентификатору для аутентифицированного пользователя
// @Tags motos
// @Produce json
// @Security BearerAuth
// @Param id path string true "Moto ID"
// @Success 200 {object} moto.Moto "Данные мотоцикла"
// @Failure 400 {object} response "Некорректный ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Мотоцикл не найден"
// @Failure 500 {object} response "Внутренняя ошибка сервера"
// @Router /api/moto/{id} [get]
func (h *MotoHandler) GetMotoById(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Некорректный moto ID")
		return
	}
	motoObj, getErr := h.service.GetMotoById(id, userId)
	if getErr != nil {
		errorResponse(c, http.StatusNotFound, getErr.Error())
		return
	}
	c.JSON(http.StatusOK, motoObj)
}

// UpdateMoto godoc
// @Summary Update a moto
// @Description Updates an existing moto for the authenticated user
// @Tags motos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Moto ID"
// @Param input body MotoRq true "Moto update data"
// @Success 200 "Moto updated successfully"
// @Failure 400 {object} response "Invalid request data or ID"
// @Failure 401 {object} response "Unauthorized"
// @Failure 404 {object} response "Moto not found"
// @Failure 500 {object} response "Internal server error"
// @Router /api/moto/{id} [put]
func (h *MotoHandler) UpdateMoto(c *gin.Context) {
	var req MotoRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid moto ID")
		return
	}

	updatedMoto := moto.NewMoto(id, userId, req.FuelType, req.Brand, req.Year, req.Plate, req.Category, req.TransmissionType)
	updatedMoto.SetVin(req.Vin)

	if err := updatedMoto.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return
	}

	updateErr := h.service.UpdateMoto(id, userId, updatedMoto)
	if updateErr != nil {
		errorResponse(c, http.StatusInternalServerError, updateErr.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteMoto godoc
// @Summary Delete a moto
// @Description Deletes an existing moto for the authenticated user
// @Tags motos
// @Produce json
// @Security BearerAuth
// @Param id path string true "Moto ID"
// @Success 200 "Moto deleted successfully"
// @Failure 400 {object} response "Invalid moto ID"
// @Failure 401 {object} response "Unauthorized"
// @Failure 404 {object} response "Moto not found"
// @Failure 500 {object} response "Internal server error"
// @Router /api/moto/{id} [delete]
func (h *MotoHandler) DeleteMoto(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid moto ID")
		return
	}
	err = h.service.DeleteMoto(id, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
