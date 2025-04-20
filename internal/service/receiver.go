package service

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/repository"
)

func ReceiveAndStoreVehicle(ch <-chan vehicle.GenericVehicle, stop <-chan struct{}) {
	go func() {
		for {
			select {
			case v := <-ch:
				repository.StoreVehicle(v)
			case <-stop:
				return
			}
		}
	}()
}
