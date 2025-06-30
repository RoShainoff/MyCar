package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/service/mock"
	"github.com/google/uuid"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	mockedId := uuid.New()
	repo := &mock.AbstractRepositoryMock{
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
		GetUsersFunc: func() []*auth.User {
			return []*auth.User{}
		},
	}
	service := NewUserService(repo)
	id, err := service.CreateUser("Ivan", "Ivanov", "ivan@example.com", "ivan1234", "Password1!")
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if id != mockedId {
		t.Errorf("ожидался id %v, получен %v", mockedId, id)
	}
}

func TestUserService_CreateUser_LoginNotFree(t *testing.T) {
	repo := &mock.AbstractRepositoryMock{
		GetUsersFunc: func() []*auth.User {
			return []*auth.User{
				{Login: "ivan1234", Id: uuid.New()},
			}
		},
	}
	service := NewUserService(repo)
	_, err := service.CreateUser("Ivan", "Ivanov", "ivan@example.com", "ivan1234", "Password1!")
	if err == nil {
		t.Error("ожидалась ошибка при занятом логине")
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	mockedId := uuid.New()
	repo := &mock.AbstractRepositoryMock{
		GetUserByIdFunc: func(id uuid.UUID) (*auth.User, *model.ApplicationError) {
			return &auth.User{Id: id, Login: "old", Password: "old"}, nil
		},
		GetUsersFunc: func() []*auth.User {
			return []*auth.User{}
		},
		SaveEntityFunc: func(entity model.BusinessEntity) (uuid.UUID, *model.ApplicationError) {
			return mockedId, nil
		},
	}
	service := NewUserService(repo)
	err := service.UpdateUser(mockedId, "Ivan", "Ivanov", "ivan@example.com", "ivan1234", "Password1!")
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}

func TestUserService_UpdateUser_LoginNotFree(t *testing.T) {
	mockedId := uuid.New()
	repo := &mock.AbstractRepositoryMock{
		GetUsersFunc: func() []*auth.User {
			return []*auth.User{
				{Login: "ivan1234", Id: uuid.New()},
			}
		},
	}
	service := NewUserService(repo)
	err := service.UpdateUser(mockedId, "Ivan", "Ivanov", "ivan@example.com", "ivan1234", "Password1!")
	if err == nil {
		t.Error("ожидалась ошибка при занятом логине")
	}
}

func TestUserService_GetUser(t *testing.T) {
	mockedId := uuid.New()
	mockUser := &auth.User{Id: mockedId, Login: "user123"}
	repo := &mock.AbstractRepositoryMock{
		GetUserByIdFunc: func(id uuid.UUID) (*auth.User, *model.ApplicationError) {
			return mockUser, nil
		},
	}
	service := NewUserService(repo)
	user, err := service.GetUser(mockedId)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if user.Id != mockedId {
		t.Errorf("ожидался id %v, получен %v", mockedId, user.Id)
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	mockedId := uuid.New()
	repo := &mock.AbstractRepositoryMock{
		GetUserByIdFunc: func(id uuid.UUID) (*auth.User, *model.ApplicationError) {
			return &auth.User{Id: id}, nil
		},
		DeleteEntityFunc: func(entity model.BusinessEntity) *model.ApplicationError {
			return nil
		},
	}
	service := NewUserService(repo)
	err := service.DeleteUser(mockedId)
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
}
