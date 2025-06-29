package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/expense"
	"MyCar/internal/repository"
	"github.com/google/uuid"
)

type AbstractExpenseService interface {
	CreateExpense(newExpense *expense.Expense, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError)
	UpdateExpense(id uuid.UUID, userId uuid.UUID, updatedExpense *expense.Expense) *model.ApplicationError
	DeleteExpense(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

type ExpenseService struct {
	repo repository.AbstractRepository
}

func NewExpenseService(repo repository.AbstractRepository) AbstractExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) CreateExpense(newExpense *expense.Expense, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return s.repo.SaveEntity(newExpense)
}

func (s *ExpenseService) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	return s.repo.GetExpenseById(id, userId)
}

func (s *ExpenseService) UpdateExpense(id uuid.UUID, userId uuid.UUID, updatedExpense *expense.Expense) *model.ApplicationError {
	expToUpdate, err := s.repo.GetExpenseById(id, userId)
	if err != nil {
		return err
	}

	expToUpdate.SetVehicleId(updatedExpense.GetVehicleId())
	expToUpdate.SetCategory(updatedExpense.GetCategory())
	expToUpdate.SetAmount(updatedExpense.GetAmount())
	expToUpdate.SetCurrency(updatedExpense.GetCurrency())
	expToUpdate.SetExchangeRate(updatedExpense.GetExchangeRate())
	expToUpdate.SetDate(updatedExpense.GetDate())
	expToUpdate.SetNote(updatedExpense.GetNote())
	expToUpdate.SetAuditFields(updatedExpense.GetAuditFields())

	_, saveErr := s.repo.SaveEntity(expToUpdate)
	return saveErr
}

func (s *ExpenseService) DeleteExpense(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	entity, err := s.repo.GetExpenseById(id, userId)
	if err != nil {
		return err
	}
	return s.repo.DeleteEntity(entity)
}
