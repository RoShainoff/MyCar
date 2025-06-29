package service

import (
	"MyCar/internal/model"
	"MyCar/internal/repository"
	"github.com/google/uuid"
)

type AbstractAttachmentService interface {
	CreateAttachment(attachment *model.Attachment, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError)
	DeleteAttachment(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

type AttachmentService struct {
	repo repository.AbstractRepository
}

func NewAttachmentService(repo repository.AbstractRepository) AbstractAttachmentService {
	return &AttachmentService{repo: repo}
}

func (s *AttachmentService) CreateAttachment(attachment *model.Attachment, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return s.repo.SaveEntity(attachment)
}

func (s *AttachmentService) GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError) {
	return s.repo.GetAttachmentById(id, userId)
}

func (s *AttachmentService) DeleteAttachment(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	entity, err := s.repo.GetAttachmentById(id, userId)
	if err != nil {
		return err
	}
	return s.repo.DeleteEntity(entity)
}
