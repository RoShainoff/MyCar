package main

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/repository"
	"MyCar/internal/service"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("System running... Press Ctrl+C to stop")
	repository.LoadAll()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	vehicleChan := make(chan vehicle.GenericVehicle)
	eventsChan := make(chan vehicle.GenericVehicle)

	service.GenerateAndSendVehicle(ctx, vehicleChan)
	service.ReceiveAndStoreVehicle(ctx, vehicleChan, eventsChan)
	service.MonitorAndLog(ctx, eventsChan)

	<-ctx.Done()
	fmt.Println("\nShutting down gracefully...")

	time.Sleep(1 * time.Second)
}
