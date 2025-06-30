package service

import (
	"MyCar/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"testing"
	"time"
)

func getTestConfig() *config.Config {
	return &config.Config{
		App: config.App{
			Secret:        "testsecret",
			TokenTtlHours: 1,
		},
	}
}

func TestJwtService_GetToken_Success(t *testing.T) {
	cfg := getTestConfig()
	svc := NewJwtService(cfg)
	id := uuid.New()
	token, err := svc.GetToken(id)
	if err != nil {
		t.Fatalf("не ожидалась ошибка, получено: %v", err)
	}
	if token == "" {
		t.Error("токен не сгенерирован")
	}
}

func TestJwtService_ParseToken_Success(t *testing.T) {
	cfg := getTestConfig()
	svc := NewJwtService(cfg)
	id := uuid.New()
	token, err := svc.GetToken(id)
	if err != nil {
		t.Fatalf("ошибка генерации токена: %v", err)
	}
	claims, err := svc.ParseToken(token)
	if err != nil {
		t.Fatalf("ошибка парсинга токена: %v", err)
	}
	if claims.UserId != id {
		t.Errorf("ожидался UserId %v, получен %v", id, claims.UserId)
	}
}

func TestJwtService_ParseToken_WrongSigningMethod(t *testing.T) {
	cfg := getTestConfig()
	svc := NewJwtService(cfg)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"exp": 9999999999})
	signed, _ := token.SignedString([]byte(cfg.App.Secret))
	_, err := svc.ParseToken(signed)
	if err == nil {
		t.Error("ожидалась ошибка при неверном методе подписи")
	}
}

func TestJwtService_ParseToken_InvalidToken(t *testing.T) {
	cfg := getTestConfig()
	svc := NewJwtService(cfg)
	// Истёкший токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(-time.Hour).Unix(),
		Issuer:    "myCar-app",
	})
	signed, _ := token.SignedString([]byte(cfg.App.Secret))
	_, err := svc.ParseToken(signed)
	if err == nil {
		t.Error("ожидалась ошибка для невалидного токена")
	}
}
