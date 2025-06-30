package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"github.com/google/uuid"
)

type AbstractRepositoryMock struct {
	SaveEntityFunc             func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError)
	DeleteEntityFunc           func(entity model.BusinessEntity) *model.ApplicationError
	GetUsersFunc               func() []*auth.User
	GetCarsFunc                func() []*car.Car
	GetMotosFunc               func() []*moto.Moto
	GetExpensesFunc            func() []*expense.Expense
	GetAttachmentsFunc         func() []*model.Attachment
	GetUserByIdFunc            func(id uuid.UUID) (*auth.User, *model.ApplicationError)
	GetCarByIdFunc             func(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError)
	GetMotoByIdFunc            func(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError)
	GetExpenseByIdFunc         func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError)
	GetAttachmentByIdFunc      func(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError)
	GetUsersCountFunc          func() int
	GetCarsCountFunc           func() int
	GetMotosCountFunc          func() int
	GetExpensesCountFunc       func() int
	GetAttachmentsCountFunc    func() int
	GetUserFunc                func(login, password string) (*auth.User, *model.ApplicationError)
	GetCarsByUserIdFunc        func(userId uuid.UUID) []*car.Car
	GetMotosByUserIdFunc       func(userId uuid.UUID) []*moto.Moto
	GetExpensesByUserIdFunc    func(userId uuid.UUID) []*expense.Expense
	GetAttachmentsByUserIdFunc func(userId uuid.UUID) []*model.Attachment
}

func (m *AbstractRepositoryMock) SaveEntity(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
	return m.SaveEntityFunc(entity)
}
func (m *AbstractRepositoryMock) DeleteEntity(entity model.BusinessEntity) *model.ApplicationError {
	return m.DeleteEntityFunc(entity)
}
func (m *AbstractRepositoryMock) GetUsers() []*auth.User {
	return m.GetUsersFunc()
}
func (m *AbstractRepositoryMock) GetCars() []*car.Car {
	return m.GetCarsFunc()
}
func (m *AbstractRepositoryMock) GetMotos() []*moto.Moto {
	return m.GetMotosFunc()
}
func (m *AbstractRepositoryMock) GetExpenses() []*expense.Expense {
	return m.GetExpensesFunc()
}
func (m *AbstractRepositoryMock) GetAttachments() []*model.Attachment {
	return m.GetAttachmentsFunc()
}
func (m *AbstractRepositoryMock) GetUserById(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	return m.GetUserByIdFunc(id)
}
func (m *AbstractRepositoryMock) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	return m.GetCarByIdFunc(id, userId)
}
func (m *AbstractRepositoryMock) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	return m.GetMotoByIdFunc(id, userId)
}
func (m *AbstractRepositoryMock) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	return m.GetExpenseByIdFunc(id, userId)
}
func (m *AbstractRepositoryMock) GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError) {
	return m.GetAttachmentByIdFunc(id, userId)
}
func (m *AbstractRepositoryMock) GetUsersCount() int {
	return m.GetUsersCountFunc()
}
func (m *AbstractRepositoryMock) GetCarsCount() int {
	return m.GetCarsCountFunc()
}
func (m *AbstractRepositoryMock) GetMotosCount() int {
	return m.GetMotosCountFunc()
}
func (m *AbstractRepositoryMock) GetExpensesCount() int {
	return m.GetExpensesCountFunc()
}
func (m *AbstractRepositoryMock) GetAttachmentsCount() int {
	return m.GetAttachmentsCountFunc()
}
func (m *AbstractRepositoryMock) GetUser(login, password string) (*auth.User, *model.ApplicationError) {
	return m.GetUserFunc(login, password)
}
func (m *AbstractRepositoryMock) GetCarsByUserId(userId uuid.UUID) []*car.Car {
	return m.GetCarsByUserIdFunc(userId)
}
func (m *AbstractRepositoryMock) GetMotosByUserId(userId uuid.UUID) []*moto.Moto {
	return m.GetMotosByUserIdFunc(userId)
}
func (m *AbstractRepositoryMock) GetExpensesByUserId(userId uuid.UUID) []*expense.Expense {
	return m.GetExpensesByUserIdFunc(userId)
}
func (m *AbstractRepositoryMock) GetAttachmentsByUserId(userId uuid.UUID) []*model.Attachment {
	return m.GetAttachmentsByUserIdFunc(userId)
}
