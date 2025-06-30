package mock

import (
	"MyCar/internal/model"
	"MyCar/internal/model/expense"
	"github.com/google/uuid"
)

type ExpenseServiceMock struct {
	CreateExpenseFunc  func(newExpense *expense.Expense, userId uuid.UUID) (uuid.UUID, *model.ApplicationError)
	GetExpenseByIdFunc func(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError)
	UpdateExpenseFunc  func(id uuid.UUID, userId uuid.UUID, updatedExpense *expense.Expense) *model.ApplicationError
	DeleteExpenseFunc  func(id uuid.UUID, userId uuid.UUID) *model.ApplicationError
}

func (m *ExpenseServiceMock) CreateExpense(newExpense *expense.Expense, userId uuid.UUID) (uuid.UUID, *model.ApplicationError) {
	return m.CreateExpenseFunc(newExpense, userId)
}
func (m *ExpenseServiceMock) GetExpenseById(id uuid.UUID, userId uuid.UUID) (*expense.Expense, *model.ApplicationError) {
	return m.GetExpenseByIdFunc(id, userId)
}
func (m *ExpenseServiceMock) UpdateExpense(id uuid.UUID, userId uuid.UUID, updatedExpense *expense.Expense) *model.ApplicationError {
	return m.UpdateExpenseFunc(id, userId, updatedExpense)
}
func (m *ExpenseServiceMock) DeleteExpense(id uuid.UUID, userId uuid.UUID) *model.ApplicationError {
	return m.DeleteExpenseFunc(id, userId)
}
