package repository

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"fmt"
)

var (
	vehicles []*vehicle.Vehicle
	cars     []*car.Car
	motos    []*moto.Moto
)

func StoreVehicle(v vehicle.GenericVehicle) {
	switch v := v.(type) {
	case *car.Car:
		cars = append(cars, v)
		fmt.Printf("Stored car: %+v\n", v.GetGeneralInfo())
	case *moto.Moto:
		motos = append(motos, v)
		fmt.Printf("Stored moto: %+v\n", v.GetGeneralInfo())
	case *vehicle.Vehicle:
		vehicles = append(vehicles, v)
		fmt.Printf("Stored vehicle: %+v\n", v.GetGeneralInfo())
	default:
		fmt.Println("Unknown vehicle type")
	}
}
