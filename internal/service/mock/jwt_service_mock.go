package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"github.com/google/uuid"
)

type JwtServiceMock struct {
	GetTokenFunc   func(id uuid.UUID) (string, *model.ApplicationError)
	ParseTokenFunc func(tokenString string) (*auth.Claims, *model.ApplicationError)
}

func (m *JwtServiceMock) GetToken(id uuid.UUID) (string, *model.ApplicationError) {
	return m.GetTokenFunc(id)
}
func (m *JwtServiceMock) ParseToken(tokenString string) (*auth.Claims, *model.ApplicationError) {
	return m.ParseTokenFunc(tokenString)
}
