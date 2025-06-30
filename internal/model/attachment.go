package model

import (
	"github.com/google/uuid"
	"time"
)

type AttachmentType string

const (
	AttachmentTypeExpense AttachmentType = "expense"
	AttachmentTypeVehicle AttachmentType = "vehicle"
)

type Attachment struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	EntityType  AttachmentType
	EntityId    uuid.UUID
	FilePath    string
	FileName    string
	MimeType    string
	UploadedAt  time.Time
	AuditFields AuditFields
}

func NewAttachment(id uuid.UUID, entityType AttachmentType, entityId uuid.UUID, filePath, fileName, mimeType string, uploadedAt time.Time, createdBy uuid.UUID) *Attachment {
	return &Attachment{
		Id:         id,
		EntityType: entityType,
		EntityId:   entityId,
		FilePath:   filePath,
		FileName:   fileName,
		MimeType:   mimeType,
		UploadedAt: uploadedAt,
		AuditFields: AuditFields{
			CreatedBy:    createdBy,
			CreatedAtUtc: time.Now(),
		},
	}
}

func (a *Attachment) GetGeneralInfo() string {
	return "Attachment: " + a.FileName
}
func (a *Attachment) GetId() uuid.UUID               { return a.Id }
func (a *Attachment) SetId(id uuid.UUID)             { a.Id = id }
func (a *Attachment) GetCreatedBy() uuid.UUID        { return a.AuditFields.GetCreatedBy() }
func (a *Attachment) SetCreatedBy(userId uuid.UUID)  { a.AuditFields.SetCreatedBy(userId) }
func (a *Attachment) GetCreatedAtUtc() time.Time     { return a.AuditFields.GetCreatedAtUtc() }
func (a *Attachment) SetCreatedAtUtc(t time.Time)    { a.AuditFields.SetCreatedAtUtc(t) }
func (a *Attachment) GetModifiedBy() uuid.UUID       { return a.AuditFields.GetModifiedBy() }
func (a *Attachment) SetModifiedBy(userId uuid.UUID) { a.AuditFields.SetModifiedBy(userId) }
func (a *Attachment) GetModifiedAtUtc() time.Time    { return a.AuditFields.GetModifiedAtUtc() }
func (a *Attachment) SetModifiedAtUtc(t time.Time)   { a.AuditFields.SetModifiedAtUtc(t) }
