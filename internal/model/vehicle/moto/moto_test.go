package moto_test

import (
	"MyCar/internal/model/vehicle"
	"MyCar/internal/model/vehicle/moto"
	"testing"
)

func TestMoto_Validate(t *testing.T) {
	tests := []struct {
		name    string
		moto    moto.Moto
		wantErr bool
	}{
		{
			name: "valid moto",
			moto: moto.Moto{
				Vehicle: vehicle.Vehicle{
					Year:        2015,
					Plate:       "MOTO123",
					FuelType:    vehicle.Diesel,
					VehicleType: vehicle.Motorcycle,
				},
				Brand:            moto.Minsk.GetBrand(),
				Category:         moto.Classic.GetCategory(),
				TransmissionType: moto.TransmissionTypeManual.GetTransmissionType(),
			},
			wantErr: false,
		},
		{
			name: "unknown brand",
			moto: moto.Moto{
				Vehicle: vehicle.Vehicle{
					Year:        2015,
					Plate:       "MOTO123",
					FuelType:    vehicle.Diesel,
					VehicleType: vehicle.Motorcycle,
				},
				Brand:            moto.UnknownBrand.GetBrand(),
				Category:         moto.Classic.GetCategory(),
				TransmissionType: moto.TransmissionTypeManual.GetTransmissionType(),
			},
			wantErr: true,
		},
		{
			name: "empty plate",
			moto: moto.Moto{
				Vehicle: vehicle.Vehicle{
					Year:        2015,
					Plate:       "",
					FuelType:    vehicle.Diesel,
					VehicleType: vehicle.Motorcycle,
				},
				Brand:            moto.Minsk.GetBrand(),
				Category:         moto.Classic.GetCategory(),
				TransmissionType: moto.TransmissionTypeManual.GetTransmissionType(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.moto.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
