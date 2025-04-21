package service

import (
	"MyCar/internal/model/vehicle"
	"math/rand"
	"time"
)

func GenerateAndSendVehicle(ch chan<- vehicle.GenericVehicle, stop <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				var v vehicle.GenericVehicle
				r := rand.Intn(3)
				if r == 0 {
					v = NewVehicle(false)
				} else if r == 1 {
					v = NewCar(false)
				} else if r == 2 {
					v = NewMoto(false)
				}
				ch <- v
			case <-stop:
				return
			}
		}
	}()
}
