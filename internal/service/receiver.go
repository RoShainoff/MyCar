package service

import (
	"MyCar/internal/model"
	"MyCar/internal/repository"
	"context"
	"fmt"
)

func ReceiveAndStoreVehicle(ctx context.Context, ch <-chan model.BusinessEntity, events chan<- model.BusinessEntity) {
	go func() {
		for {
			select {
			case be := <-ch:
				fmt.Printf("[RECV] Received entity: %s\n", be.GetGeneralInfo())

				repository.StoreGenericVehicle(be)

				events <- be
			case <-ctx.Done():
				fmt.Println("[RECV] Receiver stopped.")
				return
			}
		}
	}()
}
