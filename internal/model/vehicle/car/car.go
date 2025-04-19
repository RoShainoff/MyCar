package car

import (
	"MyCar/internal/model/vehicle"
	"fmt"
)

type Car struct {
	vehicle.Vehicle
	brand            Brand
	driveType        DriveType
	bodyType         BodyType
	transmissionType TransmissionType
}

func NewCar(id int, ownerID int, fuelType vehicle.FuelType, brand Brand, year int, plate string, driveType DriveTypeKind, bodyTypeKind BodyTypeKind, transmissionTypeKind TransmissionTypeKind) *Car {
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
	return c.brand
}

func (c *Car) GetDriveType() DriveType {
	return c.driveType
}

func (c *Car) SetDriveType(driveTypeKind DriveTypeKind) {
	c.driveType = driveTypeKind.GetDriveType()
}

func (c *Car) GetBodyType() BodyType {
	return c.bodyType
}

func (c *Car) SetBodyType(bodyTypeKind BodyTypeKind) {
	c.bodyType = bodyTypeKind.GetBodyType()
}

func (c *Car) GetTransmissionType() TransmissionType {
	return c.transmissionType
}

func (c *Car) SetTransmissionType(transmissionTypeKind TransmissionTypeKind) {
	c.transmissionType = transmissionTypeKind.GetTransmissionType()
}

func (c *Car) GetGeneralInfo() string {
	return fmt.Sprintf("This %s %s was made in %d and has VIN %s.", c.GetBodyType(), c.GetBrand(), c.GetYear(), c.GetVin())
}
