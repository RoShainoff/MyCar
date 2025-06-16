package repository

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"github.com/google/uuid"
)

type AbstractRepository interface {
	SaveEntity(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError)
	DeleteEntity(entity model.BusinessEntity) *model.ApplicationError
	GetUsers() []*auth.User
	GetCars() []*car.Car
	GetMotos() []*moto.Moto
	GetExpenses() []*expense.Expense
	GetUserById(id uuid.UUID) (*auth.User, *model.ApplicationError)
	GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError)
	GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError)
	GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError)
	GetUsersCount() int
	GetCarsCount() int
	GetMotosCount() int
	GetExpensesCount() int
	GetUser(login, password string) (*auth.User, *model.ApplicationError)
	GetCarsByUserId(userId uuid.UUID) []*car.Car
	GetMotosByUserId(userId uuid.UUID) []*moto.Moto
	GetExpensesByUserId(userId uuid.UUID) []*expense.Expense
}
