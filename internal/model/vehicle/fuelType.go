package vehicle

type FuelType int

const (
	UnknownFuelType FuelType = iota
	Petrol
	PetrolPropaneButane
	PetrolMethane
	PetrolHybrid
	Diesel
	DieselHybrid
	Electric
)
