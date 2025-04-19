package main

import (
	"MyCar/internal/service"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		service.GenerateAndStoreVehicle()
	}
}
