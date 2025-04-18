package expense

import (
	"fmt"
	"time"
)

type Expense struct {
	id           int
	vehicleId    int
	category     Category
	amount       float64
	currency     string
	exchangeRate float64
	date         time.Time
	note         string
}

func NewExpense(id, vehicleId int, category Category, amount float64, currency string, exchangeRate float64, date time.Time, note string) (*Expense, error) {
	if !isValidCategory(category) {
		return nil, fmt.Errorf("invalid expense category: %s", category)
	}
	if currency == "" {
		currency = "RUB" // по умолчанию
	}
	return &Expense{
		id:           id,
		vehicleId:    vehicleId,
		category:     category,
		amount:       amount,
		currency:     currency,
		exchangeRate: exchangeRate,
		date:         date,
		note:         note,
	}, nil
}

func (e *Expense) GetID() int {
	return e.id
}

func (e *Expense) GetVehicleID() int {
	return e.vehicleId
}

func (e *Expense) GetDate() time.Time {
	return e.date
}

func (e *Expense) GetNote() string {
	return e.note
}

func (e *Expense) GetCategory() Category {
	return e.category
}

func (e *Expense) GetAmount() float64 {
	return e.amount
}

func (e *Expense) GetCurrency() string {
	return e.currency
}

func (e *Expense) GetExchangeRate() float64 {
	return e.exchangeRate
}

func (e *Expense) GetAmountInBaseCurrency() float64 {
	if e.currency == "RUB" || e.exchangeRate == 0 {
		return e.amount
	}
	return e.amount * e.exchangeRate
}

func (e *Expense) SetCategory(cat Category) error {
	if !isValidCategory(cat) {
		return fmt.Errorf("invalid expense category: %s", cat)
	}
	e.category = cat
	return nil
}

func (e *Expense) SetCurrency(currency string) {
	e.currency = currency
}

func (e *Expense) SetExchangeRate(rate float64) {
	e.exchangeRate = rate
}
