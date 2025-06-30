package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle/moto"
	"MyCar/internal/service/mock"
	"github.com/google/uuid"
	"testing"
)

func TestMotoService_CreateMoto_Success(t *testing.T) {
	mockedId := uuid.New()
	mockRepo := &mock.AbstractRepositoryMock{
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewMotoService(mockRepo)
	newMoto := &moto.Moto{}
	id, err := service.CreateMoto(newMoto, uuid.New())
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if id != mockedId {
		t.Errorf("ожидался id %v, получен %v", mockedId, id)
	}
}

func TestMotoService_GetMotoById_Success(t *testing.T) {
	mockedId := uuid.New()
	mockMoto := &moto.Moto{}
	mockMoto.SetId(mockedId)
	mockRepo := &mock.AbstractRepositoryMock{
		GetMotoByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
			return mockMoto, nil
		},
	}
	service := NewMotoService(mockRepo)
	result, err := service.GetMotoById(mockedId, uuid.New())
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if result != mockMoto {
		t.Error("ожидался корректный объект мотоцикла")
	}
}

func TestMotoService_UpdateMoto_Success(t *testing.T) {
	mockedId := uuid.New()

	existingMoto := &moto.Moto{}
	existingMoto.SetId(mockedId)

	updatedMoto := &moto.Moto{}
	updatedMoto.SetId(mockedId)
	updatedMoto.SetBrand(moto.Minsk)

	mockRepo := &mock.AbstractRepositoryMock{
		GetMotoByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
			return existingMoto, nil
		},
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewMotoService(mockRepo)
	err := service.UpdateMoto(mockedId, uuid.New(), updatedMoto)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if existingMoto.Brand != moto.Minsk {
		t.Error("бренд не обновился")
	}
}

func TestMotoService_UpdateMoto_NotFound(t *testing.T) {
	mockRepo := &mock.AbstractRepositoryMock{
		GetMotoByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewMotoService(mockRepo)
	err := service.UpdateMoto(uuid.New(), uuid.New(), &moto.Moto{})
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии мотоцикла")
	}
}

func TestMotoService_DeleteMoto_Success(t *testing.T) {
	mockedId := uuid.New()
	mockMoto := &moto.Moto{}
	mockMoto.SetId(mockedId)
	mockRepo := &mock.AbstractRepositoryMock{
		GetMotoByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
			return mockMoto, nil
		},
		DeleteEntityFunc: func(entity model.BusinessEntity) *model.ApplicationError {
			return nil
		},
	}
	service := NewMotoService(mockRepo)
	err := service.DeleteMoto(mockedId, uuid.New())
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}

func TestMotoService_DeleteMoto_NotFound(t *testing.T) {
	mockRepo := &mock.AbstractRepositoryMock{
		GetMotoByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewMotoService(mockRepo)
	err := service.DeleteMoto(uuid.New(), uuid.New())
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии мотоцикла")
	}
}
