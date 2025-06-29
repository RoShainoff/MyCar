package auth

import (
	"MyCar/internal/model"
	"fmt"
	"github.com/google/uuid"
	"time"
	"unicode"
)

const (
	minLoginLength    = 8
	minPasswordLength = 10
)

type User struct {
	Id          uuid.UUID
	Name        string
	Surname     string
	Email       string
	Login       string
	Password    string
	AuditFields model.AuditFields
}

// --- Реализация BusinessEntity ---

func (u *User) GetGeneralInfo() string {
	return fmt.Sprintf("%s %s (%s)", u.Name, u.Surname, u.Email)
}

func (u *User) GetId() uuid.UUID {
	return u.Id
}

func (u *User) GetCreatedBy() uuid.UUID {
	return u.AuditFields.GetCreatedBy()
}

func (u *User) GetCreatedAtUtc() time.Time {
	return u.AuditFields.GetCreatedAtUtc()
}

func (u *User) GetModifiedBy() uuid.UUID {
	return u.AuditFields.GetModifiedBy()
}

func (u *User) GetModifiedAtUtc() time.Time {
	return u.AuditFields.GetModifiedAtUtc()
}

func (u *User) SetId(id uuid.UUID) {
	u.Id = id
}

func (u *User) SetCreatedBy(userId uuid.UUID) {
	u.AuditFields.SetCreatedBy(userId)
}

func (u *User) SetCreatedAtUtc(t time.Time) {
	u.AuditFields.SetCreatedAtUtc(t)
}

func (u *User) SetModifiedBy(userId uuid.UUID) {
	u.AuditFields.SetModifiedBy(userId)
}

func (u *User) SetModifiedAtUtc(t time.Time) {
	u.AuditFields.SetModifiedAtUtc(t)
}

func NewUser(name, surname, email, login, password string) (*User, *model.ApplicationError) {
	validationError := validateUser(name, surname, login, password)
	if validationError != nil {
		return nil, validationError
	}

	return &User{
		Id:       uuid.Nil,
		Name:     name,
		Surname:  surname,
		Email:    email,
		Login:    login,
		Password: password,
	}, nil
}

func (u *User) GetID() uuid.UUID {
	return u.Id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func validateUser(name string, surname string, login string, password string) *model.ApplicationError {
	personalDataValidationError := validatePersonalData(name, surname)
	if personalDataValidationError != nil {
		return personalDataValidationError
	}

	credentialsValidationError := validateCredentials(login, password)

	if credentialsValidationError != nil {
		return credentialsValidationError
	}

	return nil
}

func validatePersonalData(name string, surname string) *model.ApplicationError {
	if len(name) == 0 {
		return model.NewApplicationError(model.ErrorTypeValidation, "Имя не может быть пустым", nil)
	}

	if len(surname) == 0 {
		return model.NewApplicationError(model.ErrorTypeValidation, "Фамилия не может быть пустой", nil)
	}

	return nil
}

func validateCredentials(login string, password string) *model.ApplicationError {
	loginValidation := validateLogin(login)
	if loginValidation != nil {
		return loginValidation
	}

	passwordValidation := validatePassword(password)
	if passwordValidation != nil {
		return passwordValidation
	}

	return nil
}

func validateLogin(login string) *model.ApplicationError {
	if len(login) < minLoginLength {
		message := fmt.Sprintf("Логин слишком короткий. Пожалуйста, создайте логин длинной не меньше %d символов", minLoginLength)
		return model.NewApplicationError(model.ErrorTypeValidation, message, nil)
	}

	return nil
}

func validatePassword(password string) *model.ApplicationError {
	if len(password) < minPasswordLength {
		message := fmt.Sprintf("Пароль слишком короткий. Пожалуйста, создайте пароль длиной не менее %d символов", minPasswordLength)
		return model.NewApplicationError(model.ErrorTypeValidation, message, nil)
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return model.NewApplicationError(model.ErrorTypeValidation, "Пароль должен содержать букву верхнего регистра.", nil)
	}
	if !hasLower {
		return model.NewApplicationError(model.ErrorTypeValidation, "Пароль должен содержать букву нижнего регистра.", nil)
	}
	if !hasNumber {
		return model.NewApplicationError(model.ErrorTypeValidation, "Пароль должен содержать число.", nil)
	}
	if !hasSpecial {
		return model.NewApplicationError(model.ErrorTypeValidation, "Пароль должен содержать спецсимволы.", nil)
	}

	return nil
}

func (u *User) Validate() *model.ApplicationError {
	return validateUser(u.Name, u.Surname, u.Login, u.Password)
}
