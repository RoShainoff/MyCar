package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/repository"
	"github.com/google/uuid"
)

type AbstractCarService interface {
	CreateCar(newCar *car.Car, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError)
	UpdateCar(id uuid.UUID, userId uuid.UUID, updatedCar *car.Car) *model.ApplicationError
	DeleteCar(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

type CarService struct {
	repo repository.AbstractRepository
}

func NewCarService(repo repository.AbstractRepository) AbstractCarService {
	return &CarService{repo: repo}
}

func (s *CarService) CreateCar(newCar *car.Car, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return s.repo.SaveEntity(newCar)
}

func (s *CarService) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	return s.repo.GetCarById(id, userId)
}

func (s *CarService) UpdateCar(id uuid.UUID, userId uuid.UUID, updatedCar *car.Car) *model.ApplicationError {
	carToUpdate, err := s.repo.GetCarById(id, userId)
	if err != nil {
		return err
	}

	carToUpdate.Brand = updatedCar.Brand
	carToUpdate.SetDriveType(updatedCar.GetDriveTypeKind())
	carToUpdate.SetBodyType(updatedCar.GetBodyTypeKind())
	carToUpdate.SetTransmissionType(updatedCar.GetTransmissionTypeKind())
	carToUpdate.SetFuelType(updatedCar.GetFuelType())
	carToUpdate.SetVehicleType(updatedCar.GetVehicleType())
	carToUpdate.SetYear(updatedCar.GetYear())
	carToUpdate.Plate = updatedCar.Plate
	carToUpdate.SetUserID(updatedCar.GetUserId())
	carToUpdate.SetVin(updatedCar.GetVin())

	_, saveErr := s.repo.SaveEntity(carToUpdate)
	return saveErr
}

func (s *CarService) DeleteCar(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	entity, err := s.repo.GetCarById(id, userId)
	if err != nil {
		return err
	}
	return s.repo.DeleteEntity(entity)
}
