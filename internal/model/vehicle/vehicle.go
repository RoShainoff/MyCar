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
	id          int
	ownerId     int
	fuelType    FuelType
	vehicleType Type
	year        int
	plate       string
	vin         string
	auditFields model.AuditFields
}

func (v *Vehicle) GetId() int {
	return v.id
}

func (v *Vehicle) SetId(id int) {
	v.id = id
}

func (v *Vehicle) GetOwnerId() int {
	return v.ownerId
}

func (v *Vehicle) SetOwnerID(ownerId int) {
	v.ownerId = ownerId
}

func (v *Vehicle) GetFuelType() FuelType {
	return v.fuelType
}

func (v *Vehicle) SetFuelType(fuelType FuelType) {
	v.fuelType = fuelType
}

func (v *Vehicle) VehicleType() Type {
	return v.vehicleType
}

func (v *Vehicle) SetVehicleType(vehicleType Type) {
	v.vehicleType = vehicleType
}

func (v *Vehicle) GetYear() int {
	return v.year
}

func (v *Vehicle) SetYear(year int) {
	v.year = year
}

func (v *Vehicle) GetPlate() string {
	return v.plate
}

func (v *Vehicle) GetVin() string {
	return v.vin
}

func (v *Vehicle) SetVin(vin string) {
	v.vin = vin
}

func (v *Vehicle) GetCreatedBy() string        { return v.auditFields.GetCreatedBy() }
func (v *Vehicle) GetCreatedAtUtc() time.Time  { return v.auditFields.GetCreatedAtUtc() }
func (v *Vehicle) GetModifiedBy() string       { return v.auditFields.GetModifiedBy() }
func (v *Vehicle) GetModifiedAtUtc() time.Time { return v.auditFields.GetModifiedAtUtc() }

func (v *Vehicle) SetCreatedBy(user string)     { v.auditFields.SetCreatedBy(user) }
func (v *Vehicle) SetCreatedAtUtc(t time.Time)  { v.auditFields.SetCreatedAtUtc(t) }
func (v *Vehicle) SetModifiedBy(user string)    { v.auditFields.SetModifiedBy(user) }
func (v *Vehicle) SetModifiedAtUtc(t time.Time) { v.auditFields.SetModifiedAtUtc(t) }

func (v *Vehicle) GetGeneralInfo() string {
	return fmt.Sprintf("This vehicle was made in %d and has VIN %s", v.GetYear(), v.GetVin())
}

func NewVehicle(id, ownerID int, fuelType FuelType, vehicleType Type, year int, plate string) (*Vehicle, error) {
	return &Vehicle{
		id:          id,
		ownerId:     ownerID,
		fuelType:    fuelType,
		vehicleType: vehicleType,
		year:        year,
		plate:       plate,
	}, nil
}
