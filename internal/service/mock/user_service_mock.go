package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"github.com/google/uuid"
)

type UserServiceMock struct {
	CreateUserFunc func(name, surname, email, login, password string) (uuid.UUID, *model.ApplicationError)
	UpdateUserFunc func(id uuid.UUID, name, surname, email, login, password string) *model.ApplicationError
	GetUserFunc    func(id uuid.UUID) (*auth.User, *model.ApplicationError)
	DeleteUserFunc func(id uuid.UUID) *model.ApplicationError
}

func (m *UserServiceMock) CreateUser(name, surname, email, login, password string) (uuid.UUID, *model.ApplicationError) {
	return m.CreateUserFunc(name, surname, email, login, password)
}
func (m *UserServiceMock) UpdateUser(id uuid.UUID, name, surname, email, login, password string) *model.ApplicationError {
	return m.UpdateUserFunc(id, name, surname, email, login, password)
}
func (m *UserServiceMock) GetUser(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	return m.GetUserFunc(id)
}
func (m *UserServiceMock) DeleteUser(id uuid.UUID) *model.ApplicationError {
	return m.DeleteUserFunc(id)
}
