package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/expense"
	"MyCar/internal/service/mock"
	"github.com/google/uuid"
	"testing"
)

func TestExpenseService_CreateExpense(t *testing.T) {
	mockedId := uuid.New()
	repo := &mock.AbstractRepositoryMock{
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewExpenseService(repo)
	exp := &expense.Expense{}
	id, err := service.CreateExpense(exp, uuid.New())
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if id != mockedId {
		t.Errorf("ожидался id %v, получен %v", mockedId, id)
	}
}

func TestExpenseService_GetExpenseById(t *testing.T) {
	mockedId := uuid.New()
	mockedUserId := uuid.New()
	exp := &expense.Expense{}
	repo := &mock.AbstractRepositoryMock{
		GetExpenseByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
			if id == mockedId && userId == mockedUserId {
				return exp, nil
			}
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewExpenseService(repo)
	got, err := service.GetExpenseById(mockedId, mockedUserId)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if got != exp {
		t.Error("ожидался корректный expense")
	}
}

func TestExpenseService_UpdateExpense(t *testing.T) {
	mockedId := uuid.New()
	mockedUserId := uuid.New()
	oldExp := &expense.Expense{}
	newExp := &expense.Expense{}
	repo := &mock.AbstractRepositoryMock{
		GetExpenseByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
			return oldExp, nil
		},
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewExpenseService(repo)
	err := service.UpdateExpense(mockedId, mockedUserId, newExp)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}

func TestExpenseService_DeleteExpense(t *testing.T) {
	mockedId := uuid.New()
	mockedUserId := uuid.New()
	exp := &expense.Expense{}
	repo := &mock.AbstractRepositoryMock{
		GetExpenseByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
			return exp, nil
		},
		DeleteEntityFunc: func(entity model.BusinessEntity) *model.ApplicationError {
			return nil
		},
	}
	service := NewExpenseService(repo)
	err := service.DeleteExpense(mockedId, mockedUserId)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}

func TestExpenseService_UpdateExpense_NotFound(t *testing.T) {
	repo := &mock.AbstractRepositoryMock{
		GetExpenseByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewExpenseService(repo)
	err := service.UpdateExpense(uuid.New(), uuid.New(), &expense.Expense{})
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии расхода")
	}
}

func TestExpenseService_DeleteExpense_NotFound(t *testing.T) {
	repo := &mock.AbstractRepositoryMock{
		GetExpenseByIdFunc: func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewExpenseService(repo)
	err := service.DeleteExpense(uuid.New(), uuid.New())
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии расхода")
	}
}
