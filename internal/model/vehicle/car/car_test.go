package car_test

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/car"
	"testing"
)

func TestCar_Validate(t *testing.T) {
	tests := []struct {
		name    string
		car     car.Car
		wantErr bool
	}{
		{
			name: "valid car",
			car: car.Car{
				Vehicle: vehicle.Vehicle{
					Year:        2020,
					Plate:       "A123BC",
					FuelType:    vehicle.Diesel,
					VehicleType: vehicle.Car,
				},
				Brand:            car.AlfaRomeo.GetBrand(),
				DriveType:        car.FWD.GetDriveType(),
				BodyType:         car.Sedan.GetBodyType(),
				TransmissionType: car.TransmissionTypeManual.GetTransmissionType(),
			},
			wantErr: false,
		},
		{
			name: "empty plate",
			car: car.Car{
				Vehicle: vehicle.Vehicle{
					Year:        2020,
					Plate:       "",
					FuelType:    vehicle.Diesel,
					VehicleType: vehicle.Car,
				},
				Brand:            car.AlfaRomeo.GetBrand(),
				DriveType:        car.FWD.GetDriveType(),
				BodyType:         car.Sedan.GetBodyType(),
				TransmissionType: car.TransmissionTypeManual.GetTransmissionType(),
			},
			wantErr: true,
		},
		{
			name: "unknown brand",
			car: car.Car{
				Vehicle: vehicle.Vehicle{
					Year:        2020,
					Plate:       "A123BC",
					FuelType:    vehicle.Diesel,
					VehicleType: vehicle.Car,
				},
				Brand:            car.UnknownBrand.GetBrand(),
				DriveType:        car.FWD.GetDriveType(),
				BodyType:         car.Sedan.GetBodyType(),
				TransmissionType: car.TransmissionTypeManual.GetTransmissionType(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.car.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
