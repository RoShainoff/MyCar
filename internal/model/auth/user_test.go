package auth_test

import (
	"MyCar/internal/model/auth"
	"testing"
)

func TestUser_Validate_AdditionalCases(t *testing.T) {
	tests := []struct {
		name    string
		user    auth.User
		wantErr bool
	}{
		{
			name: "empty name",
			user: auth.User{
				Name:     "",
				Surname:  "Иванов",
				Email:    "ivan@example.com",
				Login:    "ivan1234",
				Password: "Password1!",
			},
			wantErr: true,
		},
		{
			name: "empty surname",
			user: auth.User{
				Name:     "Иван",
				Surname:  "",
				Email:    "ivan@example.com",
				Login:    "ivan1234",
				Password: "Password1!",
			},
			wantErr: true,
		},
		{
			name: "login too short",
			user: auth.User{
				Name:     "Иван",
				Surname:  "Иванов",
				Email:    "ivan@example.com",
				Login:    "ivan",
				Password: "Password1!",
			},
			wantErr: true,
		},
		{
			name: "password missing uppercase",
			user: auth.User{
				Name:     "Иван",
				Surname:  "Иванов",
				Email:    "ivan@example.com",
				Login:    "ivan1234",
				Password: "password1!",
			},
			wantErr: true,
		},
		{
			name: "password missing lowercase",
			user: auth.User{
				Name:     "Иван",
				Surname:  "Иванов",
				Email:    "ivan@example.com",
				Login:    "ivan1234",
				Password: "PASSWORD1!",
			},
			wantErr: true,
		},
		{
			name: "password missing number",
			user: auth.User{
				Name:     "Иван",
				Surname:  "Иванов",
				Email:    "ivan@example.com",
				Login:    "ivan1234",
				Password: "Password!!",
			},
			wantErr: true,
		},
		{
			name: "password missing special",
			user: auth.User{
				Name:     "Иван",
				Surname:  "Иванов",
				Email:    "ivan@example.com",
				Login:    "ivan1234",
				Password: "Password12",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
