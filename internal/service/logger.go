package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"context"
	"fmt"
)

func MonitorAndLog(ctx context.Context, events <-chan model.BusinessEntity) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[LOG] Monitor stopped.")
				return
			case be := <-events:
				switch be := be.(type) {
				case *vehicle.Vehicle:
					fmt.Printf("[LOG] New vehicle: %s\n", be.GetGeneralInfo())
				case *car.Car:
					fmt.Printf("[LOG] New car: %s\n", be.GetGeneralInfo())
				case *moto.Moto:
					fmt.Printf("[LOG] New moto: %s\n", be.GetGeneralInfo())
				case *expense.Expense:
					fmt.Printf("[LOG] New expense: %s\n", be.GetGeneralInfo())
				default:
					fmt.Println("Unknown type")
				}
			}
		}
	}()
}
