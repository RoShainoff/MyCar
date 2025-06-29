package repository

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"MyCar/internal/utils"
	"github.com/google/uuid"
)

const (
	usersFile       = "users.json"
	carsFile        = "cars.json"
	motosFile       = "motos.json"
	expensesFile    = "expenses.json"
	attachmentsFile = "attachments.json"
)

type PlainRepository struct {
	users       *BusinessEntityStorage[*auth.User]
	cars        *BusinessEntityStorage[*car.Car]
	motos       *BusinessEntityStorage[*moto.Moto]
	expenses    *BusinessEntityStorage[*expense.Expense]
	attachments *BusinessEntityStorage[*model.Attachment]
}

func NewPlainRepository() *PlainRepository {
	return &PlainRepository{
		users:       NewBusinessEntityStorage[*auth.User](usersFile),
		cars:        NewBusinessEntityStorage[*car.Car](carsFile),
		motos:       NewBusinessEntityStorage[*moto.Moto](motosFile),
		expenses:    NewBusinessEntityStorage[*expense.Expense](expensesFile),
		attachments: NewBusinessEntityStorage[*model.Attachment](attachmentsFile),
	}
}

func (r *PlainRepository) LoadAll() {
	r.users.Load()
	r.cars.Load()
	r.motos.Load()
	r.expenses.Load()
	r.attachments.Load()
}

func (r *PlainRepository) SaveEntity(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
	switch e := entity.(type) {
	case *car.Car:
		return r.cars.Save(e)
	case *moto.Moto:
		return r.motos.Save(e)
	case *expense.Expense:
		return r.expenses.Save(e)
	case *auth.User:
		return r.users.Save(e)
	case *model.Attachment:
		return r.attachments.Save(e)
	}
	return uuid.Nil, nil
}

func (r *PlainRepository) DeleteEntity(entity model.BusinessEntity) *model.ApplicationError {
	switch e := entity.(type) {
	case *car.Car:
		return r.cars.Delete(e, true)
	case *moto.Moto:
		return r.motos.Delete(e, true)
	case *expense.Expense:
		return r.expenses.Delete(e, true)
	case *auth.User:
		return r.users.Delete(e, true)
	case *model.Attachment:
		return r.attachments.Delete(e, true)
	}
	return nil
}

func (r *PlainRepository) GetUsers() []*auth.User {
	return r.users.GetAll()
}

func (r *PlainRepository) GetCars() []*car.Car {
	return r.cars.GetAll()
}

func (r *PlainRepository) GetMotos() []*moto.Moto {
	return r.motos.GetAll()
}

func (r *PlainRepository) GetExpenses() []*expense.Expense {
	return r.expenses.GetAll()
}

func (r *PlainRepository) GetAttachments() []*model.Attachment {
	return r.attachments.GetAll()
}

func (r *PlainRepository) GetUserById(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	user, err := r.users.GetById(id)
	return user, err
}

func (r *PlainRepository) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	car, err := r.cars.GetById(id)
	return car, err
}

func (r *PlainRepository) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	moto, err := r.motos.GetById(id)
	return moto, err
}

func (r *PlainRepository) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	return r.expenses.GetById(id)
}

func (r *PlainRepository) GetUsersCount() int {
	return len(r.users.GetAll())
}

func (r *PlainRepository) GetCarsCount() int {
	return len(r.cars.GetAll())
}

func (r *PlainRepository) GetMotosCount() int {
	return len(r.motos.GetAll())
}

func (r *PlainRepository) GetExpensesCount() int {
	return len(r.expenses.GetAll())
}

func (r *PlainRepository) GetAttachmentsCount() int {
	return len(r.expenses.GetAll())
}

func (r *PlainRepository) GetUser(login, password string) (*auth.User, *model.ApplicationError) {
	users := r.users.GetAll()
	for _, user := range users {
		if user.Login == login {
			arePasswordsEqual, err := utils.CompareHashAndPassword(user.Password, password)
			if err != nil {
				return nil, err
			}
			if arePasswordsEqual {
				return user, nil
			}
		}
	}
	return nil, model.NewApplicationError(model.ErrorTypeNotFound, "Пользователь не найден", nil)
}

func (r *PlainRepository) GetCarsByUserId(userId uuid.UUID) []*car.Car {
	var result []*car.Car
	for _, c := range r.cars.GetAll() {
		if c.Vehicle.GetUserId() == userId {
			result = append(result, c)
		}
	}
	return result
}

func (r *PlainRepository) GetMotosByUserId(userId uuid.UUID) []*moto.Moto {
	var result []*moto.Moto
	for _, m := range r.motos.GetAll() {
		if m.Vehicle.GetUserId() == userId {
			result = append(result, m)
		}
	}
	return result
}

func (r *PlainRepository) GetExpensesByUserId(userId uuid.UUID) []*expense.Expense {
	var result []*expense.Expense
	for _, e := range r.expenses.GetAll() {
		if e.GetCreatedBy() == userId {
			result = append(result, e)
		}
	}
	return result
}

func (r *PlainRepository) GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError) {
	att, err := r.attachments.GetById(id)
	if err != nil {
		return nil, err
	}
	if att.GetCreatedBy() != userId {
		return nil, model.NewApplicationError(model.ErrorTypeNotFound, "Вложение не найдено", nil)
	}
	return att, nil
}

func (r *PlainRepository) GetAttachmentsByUserId(userId uuid.UUID) []*model.Attachment {
	var result []*model.Attachment
	for _, a := range r.attachments.GetAll() {
		if a.GetCreatedBy() == userId {
			result = append(result, a)
		}
	}
	return result
}
