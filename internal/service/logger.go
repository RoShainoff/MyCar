package service

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"context"
	"fmt"
)

func MonitorAndLog(ctx context.Context, events <-chan vehicle.GenericVehicle) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[LOG] Monitor stopped.")
				return
			case v := <-events:
				switch v := v.(type) {
				case *vehicle.Vehicle:
					fmt.Printf("[LOG] New vehicle: %s\n", v.GetGeneralInfo())
				case *car.Car:
					fmt.Printf("[LOG] New car: %s\n", v.GetGeneralInfo())
				case *moto.Moto:
					fmt.Printf("[LOG] New moto: %s\n", v.GetGeneralInfo())
				default:
					fmt.Println("Unknown vehicle type")
				}
			}
		}
	}()
}
