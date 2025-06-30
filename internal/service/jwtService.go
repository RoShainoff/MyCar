package service

import (
	"MyCar/config"
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type JwtService interface {
	GetToken(id uuid.UUID) (string, *model.ApplicationError)
	ParseToken(tokenString string) (*auth.Claims, *model.ApplicationError)
}

type jwtService struct {
	cfg *config.Config
}

func NewJwtService(cfg *config.Config) JwtService {
	return &jwtService{
		cfg: cfg,
	}
}

func (j *jwtService) GetToken(id uuid.UUID) (string, *model.ApplicationError) {
	expirationTime := time.Now().Add(time.Duration(j.cfg.App.TokenTtlHours) * time.Hour)
	claims := auth.Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "myCar-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.cfg.App.Secret))
	if err != nil {
		return "", model.NewApplicationError(model.ErrorTypeInternal, "Ошибка при формировании токена", err)
	}
	return signedToken, nil
}

func (j *jwtService) ParseToken(tokenString string) (*auth.Claims, *model.ApplicationError) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.NewApplicationError(model.ErrorTypeAuth, "Неверный метод подписи", nil)
		}
		return []byte(j.cfg.App.Secret), nil
	})
	if err != nil || !token.Valid {
		return nil, model.NewApplicationError(model.ErrorTypeAuth, "Невалидный токен", nil)
	}
	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, model.NewApplicationError(model.ErrorTypeAuth, "Невалидный токен", nil)
}
