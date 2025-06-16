package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claims struct {
	UserId uuid.UUID
	jwt.StandardClaims
}
