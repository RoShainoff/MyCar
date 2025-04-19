package service

import (
	"MyCar/internal/repository"
	"math/rand"
)

func GenerateAndStoreVehicle() {
	r := rand.Intn(3)
	if r == 0 {
		repository.StoreVehicle(NewVehicle())
	} else if r == 1 {
		repository.StoreVehicle(NewCar())
	} else if r == 2 {
		repository.StoreVehicle(NewMoto())
	}
}
