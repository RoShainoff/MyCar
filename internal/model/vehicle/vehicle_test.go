package vehicle_test

import (
	"MyCar/internal/model/vehicle"
	"testing"
)

func TestVehicle_Validate(t *testing.T) {
	tests := []struct {
		name    string
		v       vehicle.Vehicle
		wantErr bool
	}{
		{
			name: "valid vehicle",
			v: vehicle.Vehicle{
				Year:        2010,
				Plate:       "1234AB-5",
				FuelType:    vehicle.Diesel,
				VehicleType: vehicle.Car,
			},
			wantErr: false,
		},
		{
			name: "year too old",
			v: vehicle.Vehicle{
				Year:        1900,
				Plate:       "1234AB-5",
				FuelType:    vehicle.Diesel,
				VehicleType: vehicle.Car,
			},
			wantErr: true,
		},
		{
			name: "empty plate",
			v: vehicle.Vehicle{
				Year:        2010,
				Plate:       "",
				FuelType:    vehicle.Diesel,
				VehicleType: vehicle.Car,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.v.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
