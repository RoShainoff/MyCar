package vehicle

import (
	"MyCar/internal/model"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Vehicle struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	FuelType    FuelType
	VehicleType Type
	Year        int
	Plate       string
	Vin         string
	AuditFields model.AuditFields
}

func (v *Vehicle) GetId() uuid.UUID                  { return v.Id }
func (v *Vehicle) GetUserId() uuid.UUID              { return v.UserId }
func (v *Vehicle) GetFuelType() FuelType             { return v.FuelType }
func (v *Vehicle) GetVehicleType() Type              { return v.VehicleType }
func (v *Vehicle) GetYear() int                      { return v.Year }
func (v *Vehicle) GetPlate() string                  { return v.Plate }
func (v *Vehicle) GetVin() string                    { return v.Vin }
func (v *Vehicle) GetAuditFields() model.AuditFields { return v.AuditFields }
func (v *Vehicle) GetCreatedBy() uuid.UUID           { return v.AuditFields.GetCreatedBy() }
func (v *Vehicle) GetCreatedAtUtc() time.Time        { return v.AuditFields.GetCreatedAtUtc() }
func (v *Vehicle) GetModifiedBy() uuid.UUID          { return v.AuditFields.GetModifiedBy() }
func (v *Vehicle) GetModifiedAtUtc() time.Time       { return v.AuditFields.GetModifiedAtUtc() }

func (v *Vehicle) SetId(id uuid.UUID)                 { v.Id = id }
func (v *Vehicle) SetUserID(userId uuid.UUID)         { v.UserId = userId }
func (v *Vehicle) SetFuelType(fuelType FuelType)      { v.FuelType = fuelType }
func (v *Vehicle) SetVehicleType(vehicleType Type)    { v.VehicleType = vehicleType }
func (v *Vehicle) SetYear(year int)                   { v.Year = year }
func (v *Vehicle) SetPlate(plate string)              { v.Plate = plate }
func (v *Vehicle) SetVin(vin string)                  { v.Vin = vin }
func (v *Vehicle) SetAuditFields(a model.AuditFields) { v.AuditFields = a }
func (v *Vehicle) SetCreatedBy(userId uuid.UUID)      { v.AuditFields.SetCreatedBy(userId) }
func (v *Vehicle) SetCreatedAtUtc(t time.Time)        { v.AuditFields.SetCreatedAtUtc(t) }
func (v *Vehicle) SetModifiedBy(userId uuid.UUID)     { v.AuditFields.SetModifiedBy(userId) }
func (v *Vehicle) SetModifiedAtUtc(t time.Time)       { v.AuditFields.SetModifiedAtUtc(t) }

func (v *Vehicle) GetGeneralInfo() string {
	return fmt.Sprintf("This vehicle was made in %d and has VIN %s", v.GetYear(), v.GetVin())
}

func NewVehicle(id uuid.UUID, userId uuid.UUID, fuelType FuelType, vehicleType Type, year int, plate string) (*Vehicle, error) {
	vehicle := &Vehicle{
		Id:          id,
		UserId:      userId,
		FuelType:    fuelType,
		VehicleType: vehicleType,
		Year:        year,
		Plate:       plate,
	}

	vehicle.SetCreatedBy(userId)
	vehicle.SetCreatedAtUtc(time.Now())

	return vehicle, nil
}
