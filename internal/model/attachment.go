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
	id          uuid.UUID
	entityType  AttachmentType
	entityId    uuid.UUID
	filePath    string
	fileName    string
	mimeType    string
	uploadedAt  time.Time
	AuditFields AuditFields
}

func NewAttachment(id uuid.UUID, entityType AttachmentType, entityId uuid.UUID, filePath, fileName, mimeType string, uploadedAt time.Time, createdBy uuid.UUID) *Attachment {
	return &Attachment{
		id:         id,
		entityType: entityType,
		entityId:   entityId,
		filePath:   filePath,
		fileName:   fileName,
		mimeType:   mimeType,
		uploadedAt: uploadedAt,
		AuditFields: AuditFields{
			CreatedBy:    createdBy,
			CreatedAtUtc: time.Now(),
		},
	}
}

func (a *Attachment) GetGeneralInfo() string {
	return "Attachment: " + a.fileName
}
func (a *Attachment) GetId() uuid.UUID               { return a.id }
func (a *Attachment) SetId(id uuid.UUID)             { a.id = id }
func (a *Attachment) GetCreatedBy() uuid.UUID        { return a.AuditFields.GetCreatedBy() }
func (a *Attachment) SetCreatedBy(userId uuid.UUID)  { a.AuditFields.SetCreatedBy(userId) }
func (a *Attachment) GetCreatedAtUtc() time.Time     { return a.AuditFields.GetCreatedAtUtc() }
func (a *Attachment) SetCreatedAtUtc(t time.Time)    { a.AuditFields.SetCreatedAtUtc(t) }
func (a *Attachment) GetModifiedBy() uuid.UUID       { return a.AuditFields.GetModifiedBy() }
func (a *Attachment) SetModifiedBy(userId uuid.UUID) { a.AuditFields.SetModifiedBy(userId) }
func (a *Attachment) GetModifiedAtUtc() time.Time    { return a.AuditFields.GetModifiedAtUtc() }
func (a *Attachment) SetModifiedAtUtc(t time.Time)   { a.AuditFields.SetModifiedAtUtc(t) }
