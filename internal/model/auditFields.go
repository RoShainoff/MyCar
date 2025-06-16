package model

import (
	"github.com/google/uuid"
	"time"
)

type Auditable interface {
	GetCreatedBy() uuid.UUID
	GetCreatedAtUtc() time.Time
	GetModifiedBy() uuid.UUID
	GetModifiedAtUtc() time.Time

	SetCreatedBy(userId uuid.UUID)
	SetCreatedAtUtc(t time.Time)
	SetModifiedBy(userId uuid.UUID)
	SetModifiedAtUtc(t time.Time)
}

type AuditFields struct {
	CreatedBy     uuid.UUID
	CreatedAtUtc  time.Time
	ModifiedBy    uuid.UUID
	ModifiedAtUtc time.Time
}

func (a *AuditFields) GetCreatedBy() uuid.UUID     { return a.CreatedBy }
func (a *AuditFields) GetCreatedAtUtc() time.Time  { return a.CreatedAtUtc }
func (a *AuditFields) GetModifiedBy() uuid.UUID    { return a.ModifiedBy }
func (a *AuditFields) GetModifiedAtUtc() time.Time { return a.ModifiedAtUtc }

func (a *AuditFields) SetCreatedBy(userId uuid.UUID)  { a.CreatedBy = userId }
func (a *AuditFields) SetCreatedAtUtc(t time.Time)    { a.CreatedAtUtc = t }
func (a *AuditFields) SetModifiedBy(userId uuid.UUID) { a.ModifiedBy = userId }
func (a *AuditFields) SetModifiedAtUtc(t time.Time)   { a.ModifiedAtUtc = t }
