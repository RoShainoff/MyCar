package service

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/repository"
	"context"
	"fmt"
)

func ReceiveAndStoreVehicle(ctx context.Context, ch chan vehicle.GenericVehicle) {
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Printf("[RECV] Received vehicle: %s\n", v.GetGeneralInfo())

				repository.StoreGenericVehicle(v)
			case <-ctx.Done():
				fmt.Println("[RECV] Receiver stopped.")
				return
			}
		}
	}()
}
