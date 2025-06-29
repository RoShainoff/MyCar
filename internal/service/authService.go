package service

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/repository"
)

type AbstractAuthService interface {
	AuthUser(login, password string) (string, *model.ApplicationError)
	ValidateToken(token string) (*auth.Claims, *model.ApplicationError)
}

type AuthService struct {
	repo repository.AbstractRepository
	jwt  JwtService
}

func NewAuthService(repository repository.AbstractRepository, jwtService JwtService) AbstractAuthService {
	return &AuthService{
		repo: repository,
		jwt:  jwtService,
	}
}

func (a *AuthService) AuthUser(login, password string) (string, *model.ApplicationError) {
	user, err := a.repo.GetUser(login, password)

	if err != nil {
		return "", err
	}

	return a.jwt.GetToken(user.Id)
}

func (a *AuthService) ValidateToken(token string) (*auth.Claims, *model.ApplicationError) {
	claims, err := a.jwt.ParseToken(token)

	if err != nil {
		return nil, err
	}

	_, errGetUser := a.repo.GetUserById(claims.UserId)

	if errGetUser != nil {
		return nil, errGetUser
	}
	return claims, nil
}
