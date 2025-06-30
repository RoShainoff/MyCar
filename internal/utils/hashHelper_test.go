package utils_test

import (
	"MyCar/internal/utils"
	"testing"
)

func TestGetHashAndCompare(t *testing.T) {
	password := "mySecret123"
	hash, err := utils.GetHash(password)
	if err != nil {
		t.Fatalf("GetHash() error: %v", err)
	}
	ok, err := utils.CompareHashAndPassword(hash, password)
	if err != nil || !ok {
		t.Errorf("CompareHashAndPassword() failed: %v", err)
	}
}

func TestCompareHashAndPassword_WrongPassword(t *testing.T) {
	password := "mySecret123"
	hash, _ := utils.GetHash(password)
	ok, err := utils.CompareHashAndPassword(hash, "wrongPassword")
	if err == nil || ok {
		t.Errorf("Expected error for wrong password, got ok=%v, err=%v", ok, err)
	}
}
