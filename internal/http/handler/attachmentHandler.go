package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type AttachmentHandler struct {
	service service.AbstractAttachmentService
}

type AttachmentRq struct {
	EntityType model.AttachmentType `json:"entity_type" binding:"required"`
	EntityId   uuid.UUID            `json:"entity_id" binding:"required"`
	FileName   string               `json:"file_name" binding:"required"`
	MimeType   string               `json:"mime_type" binding:"required"`
}

// NewAttachmentHandler создаёт новый обработчик вложений
func NewAttachmentHandler(s service.AbstractAttachmentService) AttachmentHandler {
	return AttachmentHandler{service: s}
}

// UploadAttachment godoc
// @Summary Загрузить вложение
// @Description Загружает новое вложение для пользователя
// @Tags attachments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body AttachmentRq true "Данные вложения"
// @Success 200 {object} map[string]interface{} "ID созданного вложения"
// @Failure 400 {object} response "Некорректные данные"
// @Failure 401 {object} response "Неавторизован"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/attachment [post]
func (h *AttachmentHandler) UploadAttachment(c *gin.Context) {
	var req AttachmentRq
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid request: %s", err.Error()))
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	filePath := "/uploads/" + req.FileName
	uploadedAt := time.Now()

	attachment := model.NewAttachment(
		uuid.Nil,
		req.EntityType,
		req.EntityId,
		filePath,
		req.FileName,
		req.MimeType,
		uploadedAt,
		userId,
	)
	id, err := h.service.CreateAttachment(attachment, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetAttachmentById godoc
// @Summary Получить вложение по ID
// @Description Возвращает вложение по идентификатору
// @Tags attachments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Attachment ID"
// @Success 200 {object} model.Attachment "Данные вложения"
// @Failure 400 {object} response "Некорректный ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Вложение не найдено"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/attachment/{id} [get]
func (h *AttachmentHandler) GetAttachmentById(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Некорректный attachment ID")
		return
	}
	att, getErr := h.service.GetAttachmentById(id, userId)
	if getErr != nil {
		errorResponse(c, http.StatusNotFound, getErr.Error())
		return
	}
	c.JSON(http.StatusOK, att)
}

// DeleteAttachment godoc
// @Summary Удалить вложение
// @Description Удаляет вложение по ID
// @Tags attachments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Attachment ID"
// @Success 200 "Вложение удалено"
// @Failure 400 {object} response "Некорректный ID"
// @Failure 401 {object} response "Неавторизован"
// @Failure 404 {object} response "Вложение не найдено"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/attachment/{id} [delete]
func (h *AttachmentHandler) DeleteAttachment(c *gin.Context) {
	userId := c.MustGet("UserId").(uuid.UUID)
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Некорректный attachment ID")
		return
	}
	err = h.service.DeleteAttachment(id, userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
