package expense

import (
	"MyCar/internal/model"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Expense struct {
	Id           uuid.UUID
	VehicleId    uuid.UUID
	Category     Category
	Amount       float64
	Currency     string
	ExchangeRate float64
	Date         time.Time
	Note         string
	AuditFields  model.AuditFields
}

func NewExpense(
	id uuid.UUID,
	vehicleId uuid.UUID,
	category Category,
	amount float64,
	currency string,
	exchangeRate float64,
	date time.Time,
	note string,
) *Expense {
	return &Expense{
		Id:           id,
		VehicleId:    vehicleId,
		Category:     category,
		Amount:       amount,
		Currency:     currency,
		ExchangeRate: exchangeRate,
		Date:         date,
		Note:         note,
	}
}

func (e *Expense) GetGeneralInfo() string {
	return fmt.Sprintf("Expense: %s, Amount: %.2f %s, Date: %s", e.Category, e.Amount, e.Currency, e.Date.Format("2006-01-02"))
}

func (e *Expense) GetId() uuid.UUID {
	return e.Id
}

func (e *Expense) SetId(id uuid.UUID) {
	e.Id = id
}

func (e *Expense) GetCreatedBy() uuid.UUID {
	return e.AuditFields.GetCreatedBy()
}

func (e *Expense) SetCreatedBy(userId uuid.UUID) {
	e.AuditFields.SetCreatedBy(userId)
}

func (e *Expense) GetCreatedAtUtc() time.Time {
	return e.AuditFields.GetCreatedAtUtc()
}

func (e *Expense) SetCreatedAtUtc(t time.Time) {
	e.AuditFields.SetCreatedAtUtc(t)
}

func (e *Expense) GetModifiedBy() uuid.UUID {
	return e.AuditFields.GetModifiedBy()
}

func (e *Expense) SetModifiedBy(userId uuid.UUID) {
	e.AuditFields.SetModifiedBy(userId)
}

func (e *Expense) GetModifiedAtUtc() time.Time {
	return e.AuditFields.GetModifiedAtUtc()
}

func (e *Expense) SetModifiedAtUtc(t time.Time) {
	e.AuditFields.SetModifiedAtUtc(t)
}

func (e *Expense) GetVehicleId() uuid.UUID {
	return e.VehicleId
}

func (e *Expense) SetVehicleId(vehicleId uuid.UUID) {
	e.VehicleId = vehicleId
}

func (e *Expense) GetCategory() Category {
	return e.Category
}

func (e *Expense) SetCategory(cat Category) error {
	if !isValidCategory(cat) {
		return fmt.Errorf("invalid expense category: %s", cat)
	}
	e.Category = cat
	return nil
}

func (e *Expense) GetAmount() float64 {
	return e.Amount
}

func (e *Expense) SetAmount(amount float64) {
	e.Amount = amount
}

func (e *Expense) GetCurrency() string {
	return e.Currency
}

func (e *Expense) SetCurrency(currency string) {
	e.Currency = currency
}

func (e *Expense) GetExchangeRate() float64 {
	return e.ExchangeRate
}

func (e *Expense) SetExchangeRate(rate float64) {
	e.ExchangeRate = rate
}

func (e *Expense) GetDate() time.Time {
	return e.Date
}

func (e *Expense) SetDate(date time.Time) {
	e.Date = date
}

func (e *Expense) GetNote() string {
	return e.Note
}

func (e *Expense) SetNote(note string) {
	e.Note = note
}

func (e *Expense) GetAuditFields() model.AuditFields {
	return e.AuditFields
}

func (e *Expense) SetAuditFields(a model.AuditFields) {
	e.AuditFields = a
}
