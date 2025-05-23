package car

import (
	"MyCar/internal/model/vehicle"
	"fmt"
	"github.com/google/uuid"
)

type Car struct {
	vehicle.Vehicle
	Brand            Brand
	DriveType        DriveType
	BodyType         BodyType
	TransmissionType TransmissionType
}

func NewCar(id uuid.UUID, ownerID int, fuelType vehicle.FuelType, brand Brand, year int, plate string, driveType DriveTypeKind, bodyTypeKind BodyTypeKind, transmissionTypeKind TransmissionTypeKind) *Car {
	newVehicle, _ := vehicle.NewVehicle(id, ownerID, fuelType, vehicle.Car, year, plate)

	return &Car{
		*newVehicle,
		brand,
		driveType.GetDriveType(),
		bodyTypeKind.GetBodyType(),
		transmissionTypeKind.GetTransmissionType(),
	}
}

func (c *Car) GetBrand() Brand {
	return c.Brand
}

func (c *Car) GetDriveType() DriveType {
	return c.DriveType
}

func (c *Car) SetDriveType(driveTypeKind DriveTypeKind) {
	c.DriveType = driveTypeKind.GetDriveType()
}

func (c *Car) GetBodyType() BodyType {
	return c.BodyType
}

func (c *Car) SetBodyType(bodyTypeKind BodyTypeKind) {
	c.BodyType = bodyTypeKind.GetBodyType()
}

func (c *Car) GetTransmissionType() TransmissionType {
	return c.TransmissionType
}

func (c *Car) SetTransmissionType(transmissionTypeKind TransmissionTypeKind) {
	c.TransmissionType = transmissionTypeKind.GetTransmissionType()
}

func (c *Car) GetGeneralInfo() string {
	return fmt.Sprintf("This %s %s was made in %d and has VIN %s.", c.GetBodyType(), c.GetBrand(), c.GetYear(), c.GetVin())
}
