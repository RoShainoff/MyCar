package model

import "time"

type Auditable interface {
	GetCreatedBy() string
	GetCreatedAtUtc() time.Time
	GetModifiedBy() string
	GetModifiedAtUtc() time.Time

	SetCreatedBy(user string)
	SetCreatedAtUtc(t time.Time)
	SetModifiedBy(user string)
	SetModifiedAtUtc(t time.Time)
}

type AuditFields struct {
	createdBy     string
	createdAtUtc  time.Time
	modifiedBy    string
	modifiedAtUtc time.Time
}

func (a *AuditFields) GetCreatedBy() string        { return a.createdBy }
func (a *AuditFields) GetCreatedAtUtc() time.Time  { return a.createdAtUtc }
func (a *AuditFields) GetModifiedBy() string       { return a.modifiedBy }
func (a *AuditFields) GetModifiedAtUtc() time.Time { return a.modifiedAtUtc }

func (a *AuditFields) SetCreatedBy(user string)     { a.createdBy = user }
func (a *AuditFields) SetCreatedAtUtc(t time.Time)  { a.createdAtUtc = t }
func (a *AuditFields) SetModifiedBy(user string)    { a.modifiedBy = user }
func (a *AuditFields) SetModifiedAtUtc(t time.Time) { a.modifiedAtUtc = t }
