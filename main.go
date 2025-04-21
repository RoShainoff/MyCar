package main

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/service"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	stop := make(chan struct{})
	vehicleChan := make(chan vehicle.GenericVehicle)

	service.GenerateAndSendVehicle(vehicleChan, stop)
	service.ReceiveAndStoreVehicle(vehicleChan, stop)
	service.MonitorAndLog(stop)

	fmt.Println("System running... Press Ctrl+C to stop")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Stopping...")

	close(stop)
	time.Sleep(1 * time.Second)
}
