package service

import (
	"context"
	"fmt"
	"time"

	"MyCar/internal/repository"
)

func MonitorAndLog(ctx context.Context) {
	go func() {
		// Получаем текущие размеры — считаем их уже "залогированными"
		lastVehicleCount := repository.GetVehiclesCount()
		lastCarCount := repository.GetCarCount()
		lastMotoCount := repository.GetMotoCount()

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				currentVehicles := repository.GetVehicles()
				currentCars := repository.GetCars()
				currentMotos := repository.GetMotos()

				if len(currentVehicles) > lastVehicleCount {
					fmt.Println("\n[LOG] New vehicles:")
					for _, v := range currentVehicles[lastVehicleCount:] {
						fmt.Printf("- %s\n", v.GetGeneralInfo())
					}
					lastVehicleCount = len(currentVehicles)
				}

				if len(currentCars) > lastCarCount {
					fmt.Println("\n[LOG] New cars:")
					for _, c := range currentCars[lastCarCount:] {
						fmt.Printf("- %s\n", c.GetGeneralInfo())
					}
					lastCarCount = len(currentCars)
				}

				if len(currentMotos) > lastMotoCount {
					fmt.Println("\n[LOG] New motos:")
					for _, m := range currentMotos[lastMotoCount:] {
						fmt.Printf("- %s\n", m.GetGeneralInfo())
					}
					lastMotoCount = len(currentMotos)
				}

			case <-ctx.Done():
				fmt.Println("[LOG] Monitor stopped.")
				return
			}
		}
	}()
}
