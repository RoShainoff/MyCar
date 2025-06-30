package expense_test

import (
	"MyCar/internal/model/expense"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestExpense_Validate(t *testing.T) {
	tests := []struct {
		name    string
		e       expense.Expense
		wantErr bool
	}{
		{
			name: "valid expense",
			e: expense.Expense{
				VehicleId: uuid.New(),
				Category:  expense.Fuel.GetCategory(),
				Amount:    100,
				Currency:  "RUB",
				Date:      time.Now(),
			},
			wantErr: false,
		},
		{
			name: "zero amount",
			e: expense.Expense{
				VehicleId: uuid.New(),
				Category:  expense.Fuel.GetCategory(),
				Amount:    0,
				Currency:  "RUB",
				Date:      time.Now(),
			},
			wantErr: true,
		},
		{
			name: "future date",
			e: expense.Expense{
				VehicleId: uuid.New(),
				Category:  expense.Fuel.GetCategory(),
				Amount:    100,
				Currency:  "RUB",
				Date:      time.Now().Add(24 * time.Hour),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.e.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
