package handler

import (
	"MyCar/internal/model/expense"
	"MyCar/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ExpenseHandler struct {
	service service.AbstractExpenseService
}

type ExpenseRq struct {
	VehicleId    uuid.UUID        `json:"vehicle_id" binding:"required"`
	Category     expense.Category `json:"category" binding:"required"`
	Amount       float64          `json:"amount" binding:"required"`
	Currency     string           `json:"currency" binding:"required"`
	ExchangeRate float64          `json:"exchange_rate"`
	Date         string           `json:"date" binding:"required"`
	Note         string           `json:"note"`
}

func NewExpenseHandler(s service.AbstractExpenseService) ExpenseHandler {
	return ExpenseHandler{service: s}
}

// CreateExpense godoc
// @Summary Создать расход
// @Description Создаёт новый расход для пользователя
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body ExpenseRq true "Данные расхода"
// @Success 200 {object} map[string]interface{} "ID созданного расхода"
// @Failure 400 {object} response "Некорректные данные"
// @Failure 401 {object} response "Неавторизован"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/expense [post]
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var req ExpenseRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	newExpense := expense.NewExpense(uuid.Nil, req.VehicleId, req.Category, req.Amount, req.Currency, req.ExchangeRate /*parse date*/, time.Now(), req.Note)

	if err := newExpense.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return
	}

	id, err := h.service.CreateExpense(newExpense, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetExpenseById godoc
// @Summary Получить расход по ID
// @Description Возвращает расход по идентификатору
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Expense ID"
// @Success 200 {object} expense.Expense "Данные расхода"
// @Failure 400 {object} response "Некорректный ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Расход не найден"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/expense/{id} [get]
func (h *ExpenseHandler) GetExpenseById(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Некорректный expense ID")
		return
	}
	exp, getErr := h.service.GetExpenseById(id, userId)
	if getErr != nil {
		errorResponse(c, http.StatusNotFound, getErr.Error())
		return
	}
	c.JSON(http.StatusOK, exp)
}

// UpdateExpense godoc
// @Summary Обновить расход
// @Description Обновляет существующий расход
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Expense ID"
// @Param input body ExpenseRq true "Данные расхода"
// @Success 200 "Расход обновлён"
// @Failure 400 {object} response "Некорректные данные или ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Расход не найден"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/expense/{id} [put]
func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	var req ExpenseRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid expense ID")
		return
	}
	updatedExpense := expense.NewExpense(id, req.VehicleId, req.Category, req.Amount, req.Currency, req.ExchangeRate /*parse date*/, time.Now(), req.Note)

	if err := updatedExpense.Validate(); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
		return
	}

	updateErr := h.service.UpdateExpense(id, userId, updatedExpense)
	if updateErr != nil {
		errorResponse(c, http.StatusInternalServerError, updateErr.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteExpense godoc
// @Summary Удалить расход
// @Description Удаляет расход по ID
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Expense ID"
// @Success 200 "Расход удалён"
// @Failure 400 {object} response "Некорректный ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Расход не найден"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/expense/{id} [delete]
func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid expense ID")
		return
	}
	err = h.service.DeleteExpense(id, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
