package main

import (
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"fmt"
	"time"
)

func main() {
	newVehicle, err := vehicle.NewVehicle(1, 1, vehicle.Diesel, vehicle.Car, 1998, "0064EE-5")
	if err != nil {
		fmt.Println("Ошибка при создании ТС:", err)
		return
	}

	fmt.Println("Новое ТС успешно создана:")
	fmt.Printf("ID: %d\n", newVehicle.GetId())
	fmt.Printf("OwnerId: %d\n", newVehicle.GetOwnerId())
	fmt.Printf("FuelType: %d\n", newVehicle.GetFuelType())
	fmt.Printf("Year: %d\n", newVehicle.GetYear())
	fmt.Printf("Plate: %s\n", newVehicle.GetPlate())

	newCar := car.NewCar(1, 1, vehicle.Diesel, car.AlfaRomeo, 1998, "0064EE-5", car.FWD, car.Sedan, car.TransmissionTypeManual)
	fmt.Println("Новая машина успешно создана:")
	fmt.Printf("ID: %d\n", newCar.GetId())
	fmt.Printf("OwnerId: %d\n", newCar.GetOwnerId())
	fmt.Printf("FuelType: %d\n", newCar.GetFuelType())
	fmt.Printf("Brand: %s\n", newCar.GetBrand())
	fmt.Printf("Year: %d\n", newCar.GetYear())
	fmt.Printf("Plate: %s\n", newCar.GetPlate())

	newMoto := moto.NewMoto(1, 1, vehicle.Diesel, moto.Minsk, 1971, "МАИ 1974", moto.Classic)
	fmt.Println("Новый мотоцикл успешно создан:")
	fmt.Printf("ID: %d\n", newMoto.GetId())
	fmt.Printf("OwnerId: %d\n", newMoto.GetOwnerId())
	fmt.Printf("FuelType: %d\n", newMoto.GetFuelType())
	fmt.Printf("Brand: %s\n", newMoto.GetBrand())
	fmt.Printf("Year: %d\n", newMoto.GetYear())
	fmt.Printf("Plate: %s\n", newMoto.GetPlate())

	expenseID := 1
	carID := 101
	category := expense.Fuel
	currency := "RUB"
	amount := 2500.0
	exchangeRate := 1.0
	date := time.Now()
	note := "Заправка на трассе"

	newExpense, err := expense.NewExpense(expenseID, carID, category, amount, currency, exchangeRate, date, note)
	if err != nil {
		fmt.Println("Ошибка при создании траты:", err)
		return
	}

	fmt.Println("Новая трата успешно создана:")
	fmt.Printf("ID: %d\n", newExpense.GetCategory())
	fmt.Printf("CarID: %d\n", newExpense.GetVehicleID())
	fmt.Printf("Категория: %s\n", newExpense.GetCategory())
	fmt.Printf("Сумма: %.2f\n", newExpense.GetAmount())
	fmt.Printf("Дата: %s\n", newExpense.GetDate().Format("2025-01-02 15:04"))
	fmt.Printf("Заметка: %s\n", newExpense.GetNote())
}
