package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle/moto"
	"MyCar/internal/repository"
	"github.com/google/uuid"
)

type AbstractMotoService interface {
	CreateMoto(newMoto *moto.Moto, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError)
	UpdateMoto(id uuid.UUID, userId uuid.UUID, updatedMoto *moto.Moto) *model.ApplicationError
	DeleteMoto(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

type MotoService struct {
	repo repository.AbstractRepository
}

func NewMotoService(repo repository.AbstractRepository) AbstractMotoService {
	return &MotoService{repo: repo}
}

func (s *MotoService) CreateMoto(newMoto *moto.Moto, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return s.repo.SaveEntity(newMoto)
}

func (s *MotoService) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	return s.repo.GetMotoById(id, userId)
}

func (s *MotoService) UpdateMoto(id uuid.UUID, userId uuid.UUID, updatedMoto *moto.Moto) *model.ApplicationError {
	motoToUpdate, err := s.repo.GetMotoById(id, userId)
	if err != nil {
		return err
	}

	motoToUpdate.Brand = updatedMoto.Brand
	motoToUpdate.SetTransmissionType(updatedMoto.GetTransmissionTypeKind())
	motoToUpdate.SetFuelType(updatedMoto.GetFuelType())
	motoToUpdate.SetVehicleType(updatedMoto.GetVehicleType())
	motoToUpdate.SetYear(updatedMoto.GetYear())
	motoToUpdate.Plate = updatedMoto.Plate
	motoToUpdate.SetUserID(updatedMoto.GetUserId())
	motoToUpdate.SetVin(updatedMoto.GetVin())

	_, saveErr := s.repo.SaveEntity(motoToUpdate)
	return saveErr
}

func (s *MotoService) DeleteMoto(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	entity, err := s.repo.GetMotoById(id, userId)
	if err != nil {
		return err
	}
	return s.repo.DeleteEntity(entity)
}
