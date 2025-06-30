package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
)

type AuthServiceMock struct {
	AuthUserFunc      func(login, password string) (string, *model.ApplicationError)
	ValidateTokenFunc func(token string) (*auth.Claims, *model.ApplicationError)
}

func (m *AuthServiceMock) AuthUser(login, password string) (string, *model.ApplicationError) {
	return m.AuthUserFunc(login, password)
}
func (m *AuthServiceMock) ValidateToken(token string) (*auth.Claims, *model.ApplicationError) {
	return m.ValidateTokenFunc(token)
}
