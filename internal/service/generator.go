package service

import (
	"MyCar/internal/model/vehicle"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func GenerateAndSendVehicle(ctx context.Context, ch chan vehicle.GenericVehicle) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[GEN] Generator stopped.")
				return
			default:
				var v vehicle.GenericVehicle
				r := rand.Intn(3)
				if r == 0 {
					v = NewVehicle(false)
				} else if r == 1 {
					v = NewCar(false)
				} else {
					v = NewMoto(false)
				}

				ch <- v
				fmt.Println("[GEN] Sent to channel:", v.GetGeneralInfo())

				time.Sleep(1 * time.Second)
			}
		}
	}()
}
