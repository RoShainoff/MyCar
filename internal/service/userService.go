package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/repository"
	"MyCar/internal/utils"
	"github.com/google/uuid"
)

const loginAlreadyUsedMessage = "пользователь с таким логином уже добавлен"

type AbstractUserService interface {
	CreateUser(name, surname, email, login, password string) (uuid.UUID, *model.ApplicationError)
	UpdateUser(id uuid.UUID, name, surname, email, login, password string) *model.ApplicationError
	GetUser(id uuid.UUID) (*auth.User, *model.ApplicationError)
	DeleteUser(id uuid.UUID) *model.ApplicationError
}

type UserService struct {
	repo repository.AbstractRepository
}

func NewUserService(repository repository.AbstractRepository) AbstractUserService {
	return &UserService{
		repo: repository,
	}
}

func (u UserService) CreateUser(name, surname, email, login, password string) (uuid.UUID, *model.ApplicationError) {
	newUser, err := auth.NewUser(name, surname, email, login, password)

	if err != nil {
		return uuid.Nil, err
	}

	if !u.isLoginFree(newUser.Login, newUser.GetId()) {
		return uuid.Nil, model.NewApplicationError(model.ErrorTypeValidation, loginAlreadyUsedMessage, nil)
	}

	passwordHash, errHash := utils.GetHash(newUser.Password)
	if errHash != nil {
		return uuid.Nil, errHash
	}

	newUser.Password = passwordHash

	id, err := u.repo.SaveEntity(newUser)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (u UserService) UpdateUser(id uuid.UUID, name, surname, email, login, password string) *model.ApplicationError {
	newUser, err := auth.NewUser(name, surname, email, login, password)

	if err != nil {
		return err
	}

	if !u.isLoginFree(newUser.Login, id) {
		return model.NewApplicationError(model.ErrorTypeValidation, loginAlreadyUsedMessage, nil)
	}

	userDb, err := u.repo.GetUserById(id)
	if err != nil {
		return err
	}

	passwordHash, errHash := utils.GetHash(newUser.Password)
	if errHash != nil {
		return errHash
	}

	userDb.Login = login
	userDb.Password = passwordHash
	userDb.Name = name
	userDb.Surname = surname

	_, err = u.repo.SaveEntity(userDb)
	if err != nil {
		return err
	}

	return nil
}

func (u UserService) GetUser(id uuid.UUID) (*auth.User, *model.ApplicationError) {
	user, err := u.repo.GetUserById(id)
	return user, err
}

func (u UserService) DeleteUser(id uuid.UUID) *model.ApplicationError {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return err
	}
	return u.repo.DeleteEntity(user)
}

func (u UserService) isLoginFree(login string, userId uuid.UUID) bool {
	users := u.repo.GetUsers()

	for _, user := range users {
		if login == user.Login && userId != user.GetId() {
			return false
		}
	}

	return true
}
