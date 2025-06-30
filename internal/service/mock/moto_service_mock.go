package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle/moto"
	"github.com/google/uuid"
)

type MotoServiceMock struct {
	CreateMotoFunc  func(newMoto *moto.Moto, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetMotoByIdFunc func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError)
	UpdateMotoFunc  func(id uuid.UUID, userId uuid.UUID, updatedMoto *moto.Moto) *model.ApplicationError
	DeleteMotoFunc  func(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

func (m *MotoServiceMock) CreateMoto(newMoto *moto.Moto, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return m.CreateMotoFunc(newMoto, userId)
}
func (m *MotoServiceMock) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	return m.GetMotoByIdFunc(id, userId)
}
func (m *MotoServiceMock) UpdateMoto(id uuid.UUID, userId uuid.UUID, updatedMoto *moto.Moto) *model.ApplicationError {
	return m.UpdateMotoFunc(id, userId, updatedMoto)
}
func (m *MotoServiceMock) DeleteMoto(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	return m.DeleteMotoFunc(id, userId)
}
