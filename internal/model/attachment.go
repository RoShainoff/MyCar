package model

import (
	"time"
)

// Attachment описывает файл, который может быть привязан к любой сущности
type Attachment struct {
	id         int
	filePath   string
	fileName   string
	mimeType   string
	uploadedAt time.Time
}

func NewAttachment(id int, filePath, fileName, mimeType string, uploadedAt time.Time) *Attachment {
	return &Attachment{
		id:         id,
		filePath:   filePath,
		fileName:   fileName,
		mimeType:   mimeType,
		uploadedAt: uploadedAt,
	}
}

func (a *Attachment) GetID() int {
	return a.id
}

func (a *Attachment) GetFilePath() string {
	return a.filePath
}

func (a *Attachment) GetFileName() string {
	return a.fileName
}

func (a *Attachment) GetMimeType() string {
	return a.mimeType
}

func (a *Attachment) GetUploadedAt() time.Time {
	return a.uploadedAt
}
