package model

import (
	"github.com/google/uuid"
	"time"
)

type BusinessEntity interface {
	GetGeneralInfo() string
	GetId() uuid.UUID
	GetCreatedBy() uuid.UUID
	GetCreatedAtUtc() time.Time
	GetModifiedBy() uuid.UUID
	GetModifiedAtUtc() time.Time
	SetId(id uuid.UUID)
	SetCreatedBy(userId uuid.UUID)
	SetCreatedAtUtc(t time.Time)
	SetModifiedBy(userId uuid.UUID)
	SetModifiedAtUtc(t time.Time)
}
