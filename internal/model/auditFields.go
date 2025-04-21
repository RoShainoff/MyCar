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
	CreatedBy     string
	CreatedAtUtc  time.Time
	ModifiedBy    string
	ModifiedAtUtc time.Time
}

func (a *AuditFields) GetCreatedBy() string        { return a.CreatedBy }
func (a *AuditFields) GetCreatedAtUtc() time.Time  { return a.CreatedAtUtc }
func (a *AuditFields) GetModifiedBy() string       { return a.ModifiedBy }
func (a *AuditFields) GetModifiedAtUtc() time.Time { return a.ModifiedAtUtc }

func (a *AuditFields) SetCreatedBy(user string)     { a.CreatedBy = user }
func (a *AuditFields) SetCreatedAtUtc(t time.Time)  { a.CreatedAtUtc = t }
func (a *AuditFields) SetModifiedBy(user string)    { a.ModifiedBy = user }
func (a *AuditFields) SetModifiedAtUtc(t time.Time) { a.ModifiedAtUtc = t }
