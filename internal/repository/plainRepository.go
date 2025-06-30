package repository

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/model/expense"
	"MyCar/internal/model/vehicle/car"
	"MyCar/internal/model/vehicle/moto"
	"github.com/google/uuid"
)

type CombinedRepository struct {
	pg    AbstractRepository
	mongo AbstractRepository
}

func NewCombinedRepository(pg, mongo AbstractRepository) *CombinedRepository {
	return &CombinedRepository{pg: pg, mongo: mongo}
}

func (r *CombinedRepository) SaveEntity(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
	switch entity.(type) {
	case *model.Attachment:
		return r.mongo.SaveEntity(entity)
	default:
		return r.pg.SaveEntity(entity)
	}
}

func (r *CombinedRepository) DeleteEntity(entity model.BusinessEntity) *model.ApplicationError {
	switch entity.(type) {
	case *model.Attachment:
		return r.mongo.DeleteEntity(entity)
	default:
		return r.pg.DeleteEntity(entity)
	}
}

func (r *CombinedRepository) GetUsers() []*auth.User              { return r.pg.GetUsers() }
func (r *CombinedRepository) GetCars() []*car.Car                 { return r.pg.GetCars() }
func (r *CombinedRepository) GetMotos() []*moto.Moto              { return r.pg.GetMotos() }
func (r *CombinedRepository) GetExpenses() []*expense.Expense     { return r.pg.GetExpenses() }
func (r *CombinedRepository) GetAttachments() []*model.Attachment { return r.mongo.GetAttachments() }
func (r *CombinedRepository) GetUserById(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	return r.pg.GetUserById(id)
}
func (r *CombinedRepository) GetCarById(id uuid.UUID, userId uuid.UUID) (*car.Car, *model.ApplicationError) {
	return r.pg.GetCarById(id, userId)
}
func (r *CombinedRepository) GetMotoById(id uuid.UUID, userId uuid.UUID) (*moto.Moto, *model.ApplicationError) {
	return r.pg.GetMotoById(id, userId)
}
func (r *CombinedRepository) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	return r.pg.GetExpenseById(id, userId)
}
func (r *CombinedRepository) GetAttachmentById(id uuid.UUID, userId uuid.UUID) (*model.Attachment, *model.ApplicationError) {
	return r.mongo.GetAttachmentById(id, userId)
}
func (r *CombinedRepository) GetUsersCount() int       { return r.pg.GetUsersCount() }
func (r *CombinedRepository) GetCarsCount() int        { return r.pg.GetCarsCount() }
func (r *CombinedRepository) GetMotosCount() int       { return r.pg.GetMotosCount() }
func (r *CombinedRepository) GetExpensesCount() int    { return r.pg.GetExpensesCount() }
func (r *CombinedRepository) GetAttachmentsCount() int { return r.mongo.GetAttachmentsCount() }
func (r *CombinedRepository) GetUser(login, password string) (*auth.User, *model.ApplicationError) {
	return r.pg.GetUser(login, password)
}
func (r *CombinedRepository) GetCarsByUserId(userId uuid.UUID) []*car.Car {
	return r.pg.GetCarsByUserId(userId)
}
func (r *CombinedRepository) GetMotosByUserId(userId uuid.UUID) []*moto.Moto {
	return r.pg.GetMotosByUserId(userId)
}
func (r *CombinedRepository) GetExpensesByUserId(userId uuid.UUID) []*expense.Expense {
	return r.pg.GetExpensesByUserId(userId)
}
func (r *CombinedRepository) GetAttachmentsByUserId(userId uuid.UUID) []*model.Attachment {
	return r.mongo.GetAttachmentsByUserId(userId)
}
