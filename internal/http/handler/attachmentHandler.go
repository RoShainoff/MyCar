package handler

import (
	"MyCar/internal/model"
	"MyCar/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
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
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param entity_type formData string true "Тип сущности"
// @Param entity_id formData string true "ID сущности"
// @Param file formData file true "Файл"
// @Success 200 {object} map[string]interface{} "ID созданного вложения"
// @Failure 400 {object} response "Некорректные данные"
// @Failure 401 {object} response "Неавторизован"
// @Failure 500 {object} response "Внутренняя ошибка"
// @Router /api/attachment [post]
func (h *AttachmentHandler) UploadAttachment(c *gin.Context) {
	entityType := c.PostForm("entity_type")
	entityIdStr := c.PostForm("entity_id")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Файл обязателен")
		return
	}
	entityId, err := uuid.Parse(entityIdStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Некорректный entity_id")
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Ошибка открытия файла")
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Ошибка чтения файла")
		return
	}
	userId := c.MustGet("UserId").(uuid.UUID)
	uploadedAt := time.Now()
	attachment := model.NewAttachment(
		uuid.Nil,
		model.AttachmentType(entityType),
		entityId,
		"", // filePath не нужен, если храним в MongoDB
		fileHeader.Filename,
		fileHeader.Header.Get("Content-Type"),
		uploadedAt,
		userId,
		data,
	)
	id, errApp := h.service.CreateAttachment(attachment, userId)
	if errApp != nil {
		errorResponse(c, http.StatusInternalServerError, errApp.Error())
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
