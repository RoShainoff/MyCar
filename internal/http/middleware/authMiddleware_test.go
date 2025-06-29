package middleware

import (
	"MyCar/internal/model"
	"MyCar/internal/model/auth"
	"MyCar/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	userId := uuid.New()
	mockAuthService := &mock.AuthServiceMock{
		ValidateTokenFunc: func(token string) (*auth.Claims, *model.ApplicationError) {
			return &auth.Claims{UserId: userId}, nil
		},
	}
	mw := AuthMiddleware(mockAuthService)

	router := gin.New()
	router.Use(mw)
	router.GET("/protected", func(c *gin.Context) {
		uid, exists := c.Get("UserId")
		if !exists || uid != userId {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer validtoken")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ожидался 200, получен %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockAuthService := &mock.AuthServiceMock{
		ValidateTokenFunc: func(token string) (*auth.Claims, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeAuth, "invalid", nil)
		},
	}
	mw := AuthMiddleware(mockAuthService)

	router := gin.New()
	router.Use(mw)
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("ожидался 401, получен %d", w.Code)
	}
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockAuthService := &mock.AuthServiceMock{
		ValidateTokenFunc: func(token string) (*auth.Claims, *model.ApplicationError) {
			return nil, model.NewApplicationError(model.ErrorTypeAuth, "missing", nil)
		},
	}
	mw := AuthMiddleware(mockAuthService)

	router := gin.New()
	router.Use(mw)
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("ожидался 401, получен %d", w.Code)
	}
}
