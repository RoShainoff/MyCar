package vehicle

type Vehicle struct {
	id          int
	ownerId     int
	fuelType    FuelType
	vehicleType Type
	year        int
	plate       string
	vin         string
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
