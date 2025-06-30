package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle/car"
	"github.com/google/uuid"
)

type CarServiceMock struct {
	CreateCarFunc  func(newCar *car.Car, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetCarByIdFunc func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError)
	UpdateCarFunc  func(id uuid.UUID, userId uuid.UUID, updatedCar *car.Car) *model.ApplicationError
	DeleteCarFunc  func(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

func (m *CarServiceMock) CreateCar(newCar *car.Car, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return m.CreateCarFunc(newCar, userId)
}
func (m *CarServiceMock) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	return m.GetCarByIdFunc(id, userId)
}
func (m *CarServiceMock) UpdateCar(id uuid.UUID, userId uuid.UUID, updatedCar *car.Car) *model.ApplicationError {
	return m.UpdateCarFunc(id, userId, updatedCar)
}
func (m *CarServiceMock) DeleteCar(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	return m.DeleteCarFunc(id, userId)
}
