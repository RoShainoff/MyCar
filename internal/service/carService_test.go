package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/service/mock"
	"github.com/google/uuid"
	"testing"
)

func TestCarService_CreateCar(t *testing.T) {
	mockedId := uuid.New()
	repo := &mock.AbstractRepositoryMock{
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewCarService(repo)
	c := &car.Car{}
	id, err := service.CreateCar(c, uuid.New())
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if id != mockedId {
		t.Errorf("ожидался id %v, получен %v", mockedId, id)
	}
}

func TestCarService_GetCarById(t *testing.T) {
	mockedId := uuid.New()
	mockedUserId := uuid.New()
	c := &car.Car{}
	repo := &mock.AbstractRepositoryMock{
		GetCarByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
			if id == mockedId && userId == mockedUserId {
				return c, nil
			}
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewCarService(repo)
	got, err := service.GetCarById(mockedId, mockedUserId)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if got != c {
		t.Error("ожидалась корректная машина")
	}
}

func TestCarService_UpdateCar(t *testing.T) {
	mockedId := uuid.New()
	mockedUserId := uuid.New()
	oldCar := &car.Car{}
	newCar := &car.Car{}
	repo := &mock.AbstractRepositoryMock{
		GetCarByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
			return oldCar, nil
		},
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewCarService(repo)
	err := service.UpdateCar(mockedId, mockedUserId, newCar)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}

func TestCarService_DeleteCar(t *testing.T) {
	mockedId := uuid.New()
	mockedUserId := uuid.New()
	c := &car.Car{}
	repo := &mock.AbstractRepositoryMock{
		GetCarByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
			return c, nil
		},
		DeleteEntityFunc: func(entity model.BusinessEntity) *model.ApplicationError {
			return nil
		},
	}
	service := NewCarService(repo)
	err := service.DeleteCar(mockedId, mockedUserId)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}

func TestCarService_UpdateCar_NotFound(t *testing.T) {
	repo := &mock.AbstractRepositoryMock{
		GetCarByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewCarService(repo)
	err := service.UpdateCar(uuid.New(), uuid.New(), &car.Car{})
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии машины")
	}
}

func TestCarService_DeleteCar_NotFound(t *testing.T) {
	repo := &mock.AbstractRepositoryMock{
		GetCarByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewCarService(repo)
	err := service.DeleteCar(uuid.New(), uuid.New())
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии машины")
	}
}
