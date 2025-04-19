package moto

import (
	"MyCar/internal/model/vehicle"
	"fmt"
)

type Moto struct {
	vehicle.Vehicle
	brand    Brand
	category Category
}

func NewMoto(id, ownerID int, fuelType vehicle.FuelType, brand Brand, year int, plate string, category CategoryKind) *Moto {
	newVehicle, _ := vehicle.NewVehicle(id, ownerID, fuelType, vehicle.Motorcycle, year, plate)

	return &Moto{
		*newVehicle,
		brand,
		category.GetCategory(),
	}
}

func (m *Moto) GetCategory() Category {
	return m.category
}

func (m *Moto) GetCategoryName() string {
	return m.category.name
}

func (m *Moto) GetBrand() Brand {
	return m.brand
}

func (m *Moto) GetGeneralInfo() string {
	return fmt.Sprintf("This %s was made in %d and has VIN %s. Categorised as a %s", m.GetBrand(), m.GetYear(), m.GetVin(), m.GetCategoryName())
}
