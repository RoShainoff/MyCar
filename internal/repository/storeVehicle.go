package repository

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"sync"
)

var (
	vehicles   []*vehicle.Vehicle
	cars       []*car.Car
	motos      []*moto.Moto
	vehiclesMu sync.RWMutex
	carsMu     sync.RWMutex
	motosMu    sync.RWMutex
)

func StoreVehicle(v vehicle.GenericVehicle) {
	switch v := v.(type) {
	case *car.Car:
		carsMu.Lock()
		cars = append(cars, v)
		carsMu.Unlock()
	case *moto.Moto:
		motosMu.Lock()
		motos = append(motos, v)
		motosMu.Unlock()
	case *vehicle.Vehicle:
		vehiclesMu.Lock()
		vehicles = append(vehicles, v)
		vehiclesMu.Unlock()
	}
}

func GetVehicles() []*vehicle.Vehicle {
	carsMu.RLock()
	defer carsMu.RUnlock()
	return append([]*vehicle.Vehicle(nil), vehicles...)
}

func GetCars() []*car.Car {
	carsMu.RLock()
	defer carsMu.RUnlock()
	return append([]*car.Car(nil), cars...)
}

func GetMotos() []*moto.Moto {
	motosMu.RLock()
	defer motosMu.RUnlock()
	return append([]*moto.Moto(nil), motos...)
}
