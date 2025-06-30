package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/service/mock"
	"github.com/google/uuid"
	"testing"
)

func TestAuthService_AuthUser_Success(t *testing.T) {
	mockedId := uuid.New()
	mockRepo := &mock.AbstractRepositoryMock{
		GetUserFunc: func(login, password string) (*auth.User, *model.ApplicationError) {
			return &auth.User{Id: mockedId, Login: login}, nil
		},
	}
	mockJwt := &mock.JwtServiceMock{
		GetTokenFunc: func(id uuid.UUID) (string, *model.ApplicationError) {
			return "token123", nil
		},
	}
	service := NewAuthService(mockRepo, mockJwt)
	token, err := service.AuthUser("user", "pass")
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if token != "token123" {
		t.Errorf("ожидался токен token123, получен %v", token)
	}
}

func TestAuthService_AuthUser_InvalidCredentials(t *testing.T) {
	mockRepo := &mock.AbstractRepositoryMock{
		GetUserFunc: func(login, password string) (*auth.User, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeAuth, "invalid", nil)
		},
	}
	mockJwt := &mock.JwtServiceMock{}
	service := NewAuthService(mockRepo, mockJwt)
	_, err := service.AuthUser("user", "wrongpass")
	if err == nil {
		t.Error("ожидалась ошибка при неверных данных")
	}
}

func TestAuthService_ValidateToken_Success(t *testing.T) {
	mockedId := uuid.New()
	mockClaims := &auth.Claims{UserId: mockedId}
	mockJwt := &mock.JwtServiceMock{
		ParseTokenFunc: func(token string) (*auth.Claims, *model.ApplicationError) {
			return mockClaims, nil
		},
	}
	mockRepo := &mock.AbstractRepositoryMock{
		GetUserByIdFunc: func(id uuid.UUID) (*auth.User, *model.ApplicationError) {
			return &auth.User{Id: id}, nil
		},
	}
	service := NewAuthService(mockRepo, mockJwt)
	claims, err := service.ValidateToken("sometoken")
	if err != nil {
		t.Fatalf("не ожидалась ошибка: %v", err)
	}
	if claims.UserId != mockedId {
		t.Errorf("ожидался UserId %v, получен %v", mockedId, claims.UserId)
	}
}

func TestAuthService_ValidateToken_InvalidToken(t *testing.T) {
	mockJwt := &mock.JwtServiceMock{
		ParseTokenFunc: func(token string) (*auth.Claims, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeAuth, "invalid", nil)
		},
	}
	mockRepo := &mock.AbstractRepositoryMock{}
	service := NewAuthService(mockRepo, mockJwt)
	_, err := service.ValidateToken("badtoken")
	if err == nil {
		t.Error("ожидалась ошибка при невалидном токене")
	}
}

func TestAuthService_ValidateToken_UserNotFound(t *testing.T) {
	mockedId := uuid.New()
	mockJwt := &mock.JwtServiceMock{
		ParseTokenFunc: func(token string) (*auth.Claims, *model.ApplicationError) {
			return &auth.Claims{UserId: mockedId}, nil
		},
	}
	mockRepo := &mock.AbstractRepositoryMock{
		GetUserByIdFunc: func(id uuid.UUID) (*auth.User, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeNotFound, "not found", nil)
		},
	}
	service := NewAuthService(mockRepo, mockJwt)
	_, err := service.ValidateToken("token")
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии пользователя")
	}
}
