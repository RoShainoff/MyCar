package vehicle

import (
	"MyCar/internal/model"
	"fmt"
	"time"
)

type GenericVehicle interface {
	GetGeneralInfo() string
}

type Vehicle struct {
	Id          int
	OwnerId     int
	FuelType    FuelType
	VehicleType Type
	Year        int
	Plate       string
	Vin         string
	AuditFields model.AuditFields
}

func (v *Vehicle) GetId() int {
	return v.Id
}

func (v *Vehicle) SetId(id int) {
	v.Id = id
}

func (v *Vehicle) GetOwnerId() int {
	return v.OwnerId
}

func (v *Vehicle) SetOwnerID(ownerId int) {
	v.OwnerId = ownerId
}

func (v *Vehicle) GetFuelType() FuelType {
	return v.FuelType
}

func (v *Vehicle) SetFuelType(fuelType FuelType) {
	v.FuelType = fuelType
}

func (v *Vehicle) GetVehicleType() Type {
	return v.VehicleType
}

func (v *Vehicle) SetVehicleType(vehicleType Type) {
	v.VehicleType = vehicleType
}

func (v *Vehicle) GetYear() int {
	return v.Year
}

func (v *Vehicle) SetYear(year int) {
	v.Year = year
}

func (v *Vehicle) GetPlate() string {
	return v.Plate
}

func (v *Vehicle) GetVin() string {
	return v.Vin
}

func (v *Vehicle) SetVin(vin string) {
	v.Vin = vin
}

func (v *Vehicle) GetCreatedBy() string        { return v.AuditFields.GetCreatedBy() }
func (v *Vehicle) GetCreatedAtUtc() time.Time  { return v.AuditFields.GetCreatedAtUtc() }
func (v *Vehicle) GetModifiedBy() string       { return v.AuditFields.GetModifiedBy() }
func (v *Vehicle) GetModifiedAtUtc() time.Time { return v.AuditFields.GetModifiedAtUtc() }

func (v *Vehicle) SetCreatedBy(user string)     { v.AuditFields.SetCreatedBy(user) }
func (v *Vehicle) SetCreatedAtUtc(t time.Time)  { v.AuditFields.SetCreatedAtUtc(t) }
func (v *Vehicle) SetModifiedBy(user string)    { v.AuditFields.SetModifiedBy(user) }
func (v *Vehicle) SetModifiedAtUtc(t time.Time) { v.AuditFields.SetModifiedAtUtc(t) }

func (v *Vehicle) GetGeneralInfo() string {
	return fmt.Sprintf("This vehicle was made in %d and has VIN %s", v.GetYear(), v.GetVin())
}

func NewVehicle(id, ownerID int, fuelType FuelType, vehicleType Type, year int, plate string) (*Vehicle, error) {
	return &Vehicle{
		Id:          id,
		OwnerId:     ownerID,
		FuelType:    fuelType,
		VehicleType: vehicleType,
		Year:        year,
		Plate:       plate,
	}, nil
}
