package service

import (
	"MyCar/internal/repository"
	"fmt"
	"time"
)

func MonitorAndLog(stop <-chan struct{}) {
	go func() {
		var lastVehicleCount, lastCarCount, lastMotoCount int

		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				vehicles := repository.GetVehicles()
				cars := repository.GetCars()
				motos := repository.GetMotos()

				if len(vehicles) > lastVehicleCount {
					fmt.Println("New vehicles:")
					for _, v := range vehicles[lastVehicleCount:] {
						fmt.Printf("- %s\n", v.GetGeneralInfo())
					}
					lastVehicleCount = len(vehicles)
				}

				if len(cars) > lastCarCount {
					fmt.Println("New cars:")
					for _, c := range cars[lastCarCount:] {
						fmt.Printf("- %s\n", c.GetGeneralInfo())
					}
					lastCarCount = len(cars)
				}

				if len(motos) > lastMotoCount {
					fmt.Println("New motos:")
					for _, m := range motos[lastMotoCount:] {
						fmt.Printf("- %s\n", m.GetGeneralInfo())
					}
					lastMotoCount = len(motos)
				}

			case <-stop:
				return
			}
		}
	}()
}
