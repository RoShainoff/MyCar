package moto

import (
	"MyCar/internal/model/vehicle"
	"fmt"
	"github.com/google/uuid"
)

type Moto struct {
	vehicle.Vehicle
	Brand            Brand
	Category         Category
	TransmissionType TransmissionType
}

func NewMoto(id uuid.UUID, userId uuid.UUID, fuelType vehicle.FuelType, brand Brand, year int, plate string, category CategoryKind, transmission TransmissionTypeKind) *Moto {
	newVehicle, _ := vehicle.NewVehicle(id, userId, fuelType, vehicle.Motorcycle, year, plate)
	return &Moto{
		*newVehicle,
		brand,
		category.GetCategory(),
		TransmissionTypeKind.GetTransmissionType(transmission),
	}
}

func (m *Moto) GetCategory() Category {
	return m.Category
}

func (m *Moto) GetCategoryKind() CategoryKind {
	return m.Category.id
}

func (m *Moto) SetCategory(category CategoryKind) {
	m.Category = category.GetCategory()
}

func (m *Moto) GetBrand() Brand {
	return m.Brand
}

func (m *Moto) SetBrand(brand Brand) {
	m.Brand = brand
}

func (m *Moto) GetTransmissionTypeKind() TransmissionTypeKind {
	return m.TransmissionType.Id
}

func (m *Moto) SetTransmissionType(kind TransmissionTypeKind) {
	m.TransmissionType = TransmissionTypeKind.GetTransmissionType(kind)
}

func (m *Moto) GetFuelType() vehicle.FuelType {
	return m.Vehicle.GetFuelType()
}

func (m *Moto) SetFuelType(fuelType vehicle.FuelType) {
	m.Vehicle.SetFuelType(fuelType)
}

func (m *Moto) GetVehicleType() vehicle.Type {
	return m.Vehicle.GetVehicleType()
}

func (m *Moto) SetVehicleType(vehicleType vehicle.Type) {
	m.Vehicle.SetVehicleType(vehicleType)
}

func (m *Moto) GetYear() int {
	return m.Vehicle.GetYear()
}

func (m *Moto) SetYear(year int) {
	m.Vehicle.SetYear(year)
}

func (m *Moto) GetPlate() string {
	return m.Vehicle.GetPlate()
}

func (m *Moto) SetPlate(plate string) {
	m.Vehicle.SetPlate(plate)
}

func (m *Moto) GetUserId() uuid.UUID {
	return m.Vehicle.GetUserId()
}

func (m *Moto) SetUserID(userId uuid.UUID) {
	m.Vehicle.SetUserID(userId)
}

func (m *Moto) GetVin() string {
	return m.Vehicle.GetVin()
}

func (m *Moto) SetVin(vin string) {
	m.Vehicle.SetVin(vin)
}

func (m *Moto) GetGeneralInfo() string {
	return fmt.Sprintf("This %s was made in %d and has VIN %s. Categorised as a %s", m.GetBrand(), m.GetYear(), m.GetVin(), m.GetCategoryName())
}

func (m *Moto) GetCategoryName() string {
	return m.Category.name
}
