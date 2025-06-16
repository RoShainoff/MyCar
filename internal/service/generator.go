package service

import (
	"MyCar/internal/model"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func GenerateAndSendVehicle(ctx context.Context, ch chan model.BusinessEntity) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[GEN] Generator stopped.")
				return
			default:
				var be model.BusinessEntity
				r := rand.Intn(3)
				if r == 0 {
					be = NewVehicle(false)
				} else if r == 1 {
					be = NewCar(false)
				} else {
					be = NewMoto(false)
				}

				ch <- be
				fmt.Println("[GEN] Sent to channel:", be.GetGeneralInfo())

				time.Sleep(1 * time.Second)
			}
		}
	}()
}
