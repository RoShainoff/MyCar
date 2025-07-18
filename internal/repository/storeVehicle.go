package repository

import (
	"MyCar/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
)

var (
	vehicles []*vehicle.Vehicle
	cars     []*car.Car
	motos    []*moto.Moto

	vehiclesMu sync.RWMutex
	carsMu     sync.RWMutex
	motosMu    sync.RWMutex

	vehiclesStoreMu sync.RWMutex
	carsStoreMu     sync.RWMutex
	motosStoreMu    sync.RWMutex
)

const (
	vehiclesJSON = "vehicles.json"
	carsJSON     = "cars.json"
	motosJSON    = "motos.json"
)

func LoadAll() {
	fmt.Println("Loading data...")
	loadFromJSON(vehiclesJSON, &vehicles, &vehiclesMu)
	loadFromJSON(carsJSON, &cars, &carsMu)
	loadFromJSON(motosJSON, &motos, &motosMu)
	fmt.Println("Data loaded.")
}

func StoreVehicle(v *vehicle.Vehicle) {
	vehiclesMu.Lock()
	defer vehiclesMu.Unlock()
	vehicles = append(vehicles, v)
	appendToJSON(vehiclesJSON, v, &vehiclesStoreMu)
}

func StoreCar(c *car.Car) {
	carsMu.Lock()
	defer carsMu.Unlock()
	cars = append(cars, c)
	appendToJSON(carsJSON, c, &carsStoreMu)
}

func StoreMoto(m *moto.Moto) {
	motosMu.Lock()
	defer motosMu.Unlock()
	motos = append(motos, m)
	appendToJSON(motosJSON, m, &motosStoreMu)
}

func StoreGenericVehicle(be model.BusinessEntity) {
	switch be := be.(type) {
	case *car.Car:
		StoreCar(be)
	case *moto.Moto:
		StoreMoto(be)
	case *vehicle.Vehicle:
		StoreVehicle(be)
	}
}

// ------------------ SAVE/LOAD JSON ------------------

func appendToJSON(filename string, source any, mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()

	fileExists := false
	fileInfo, err := os.Stat(filename)
	if err == nil && fileInfo.Size() > 0 {
		fileExists = true
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	if !fileExists {
		file.WriteString("[\n\t")
	} else {
		stat, _ := file.Stat()
		if stat.Size() >= 2 {
			file.Seek(-2, io.SeekEnd)

			buf := make([]byte, 1)
			file.Read(buf)
			if buf[0] != '[' {
				file.WriteString(",\n\t")
			}
		}
	}

	data, err := json.MarshalIndent(source, "\t", "\t")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	file.Write(data)

	file.WriteString("\n]")
}

func loadFromJSON[T any](filename string, target *[]T, mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("No %s found (probably first run): %v\n", filename, err)
		return
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(target); err != nil {
		fmt.Printf("Error decoding %s: %v\n", filename, err)
	}

	fmt.Printf("Loaded %d objects from the %s\n", len(*target), filename)
}
