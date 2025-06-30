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
	Id          uuid.UUID      `bson:"id"`
	UserId      uuid.UUID      `bson:"user_id"`
	EntityType  AttachmentType `bson:"entity_type"`
	EntityId    uuid.UUID      `bson:"entity_id"`
	FilePath    string         `bson:"file_path"`
	FileName    string         `bson:"file_name"`
	MimeType    string         `bson:"mime_type"`
	UploadedAt  time.Time      `bson:"uploaded_at"`
	Data        []byte         `bson:"data"`
	AuditFields AuditFields    `bson:"auditfields"`
}

func NewAttachment(id uuid.UUID, entityType AttachmentType, entityId uuid.UUID, filePath, fileName, mimeType string, uploadedAt time.Time, createdBy uuid.UUID, data []byte) *Attachment {
	return &Attachment{
		Id:         id,
		EntityType: entityType,
		EntityId:   entityId,
		FilePath:   filePath,
		FileName:   fileName,
		MimeType:   mimeType,
		UploadedAt: uploadedAt,
		Data:       data,
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
