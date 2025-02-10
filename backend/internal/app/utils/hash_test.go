package utils

import (
	"bytes"
	"testing"
	"time"

	"todo_app_backend/config"
	"todo_app_backend/internal/app/models"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	t.Run("Test valid password", func(t *testing.T) {
		hash, err := HashPassword("password")
		if err != nil {
			t.Fatal(err)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte("password")); err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
	t.Run("Test invalid password", func(t *testing.T) {
		hash, err := HashPassword("password")
		if err != nil {
			t.Fatal(err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte("wrong_password"))
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("Test valid password", func(t *testing.T) {
		hash, _ := HashPassword("password")
		if !CheckPasswordHash("password", hash) {
			t.Errorf("Expected true, got false")
		}
	})
	t.Run("Test invalid password", func(t *testing.T) {
		hash, _ := HashPassword("password")
		if CheckPasswordHash("wrong_password", hash) {
			t.Errorf("Expected false, got true")
		}
	})
	t.Run("Test nil password", func(t *testing.T) {
		if CheckPasswordHash("password", "") {
			t.Errorf("Expected false, got true")
		}
	})
	t.Run("Test nil hash", func(t *testing.T) {
		if CheckPasswordHash("password", "") {
			t.Errorf("Expected false, got true")
		}
	})
}

func TestGenerateToken(t *testing.T) {
	config.LoadConfigurationFile("../../../.env")
	t.Run("Test valid token generation", func(t *testing.T) {
		secret := []byte("secret_key")
		cfg, err := config.GetConfig()
		if err != nil || cfg == nil {
			t.Errorf("Error getting the config : %v", err)
		}
		*cfg = config.Config{TokenSecret: secret}
		token := models.Token{Exp: time.Now().AddDate(0, 0, 10)}
		tokenStr, err := GenerateToken(token)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		tokenDetails, err := ValidateToken(tokenStr)
		if err != nil {
			t.Errorf("Expected nil, got : %v", err)
		}
		if tokenDetails.UserID != 0 {
			t.Errorf("Expected UserID 0, got %d", tokenDetails.UserID)
		}
		if tokenDetails.Exp.Before(time.Now()) {
			t.Errorf("Expected token to be valid, got expired")
		}
	})
}

func TestValidateToken(t *testing.T) {
	config.LoadConfigurationFile("../../../.env")
	t.Run("Test valid token validation", func(t *testing.T) {
		secret := []byte("secret_key")
		cfg, err := config.GetConfig()
		if err != nil || cfg == nil {
			t.Errorf("Error getting the config : %v", err)
		}
		*cfg = config.Config{TokenSecret: secret}
		token := models.Token{Exp: time.Now().AddDate(0, 0, 10)}
		tokenStr, _ := GenerateToken(token)
		tokenDetails, err := ValidateToken(tokenStr)
		if err != nil {
			t.Errorf("Expected nil, got : %v", err)
		}
		if tokenDetails.UserID != 0 {
			t.Errorf("Expected UserID 0, got %d", tokenDetails.UserID)
		}
		if tokenDetails.Exp.Before(time.Now()) {
			t.Errorf("Expected token to be valid, got expired")
		}
	})
	t.Run("Test invalid token hash", func(t *testing.T) {
		secret := []byte("secret_key")
		cfg, err := config.GetConfig()
		if err != nil || cfg == nil {
			t.Errorf("Error getting the config : %v", err)
		}
		*cfg = config.Config{TokenSecret: secret}
		token := models.Token{Exp: time.Now()}
		tokenStr, _ := GenerateToken(token)
		tokenDetails, err := ValidateToken(tokenStr)
		if err != nil {
			t.Errorf("Expected nil, got : %v", err)
		}
		if tokenDetails.UserID != 0 {
			t.Errorf("Expected UserID 0, got %d", tokenDetails.UserID)
		}
	})
	t.Run("Test expired token", func(t *testing.T) {
		secret := []byte("secret_key")
		cfg, err := config.GetConfig()
		if err != nil || cfg == nil {
			t.Errorf("Error getting the config : %v", err)
		}
		*cfg = config.Config{TokenSecret: secret}
		token := models.Token{Exp: time.Now().Add(-time.Hour)}
		tokenStr, _ := GenerateToken(token)
		tokenDetails, _ := ValidateToken(tokenStr)
		if tokenDetails != nil {
			t.Errorf("Expected nil, got : %v", tokenDetails)
		}
	})
}

func TestConfig(t *testing.T) {
	config.LoadConfigurationFile("../../../.env")
	t.Run("Test SetConfig method", func(t *testing.T) {
		secret := []byte("secret_key")
		cfg, err := config.GetConfig()
		if err != nil || cfg == nil {
			t.Errorf("Error getting the config : %v", err)
		}
		*cfg = config.Config{TokenSecret: secret}
		if !bytes.Equal(cfg.TokenSecret, secret) {
			t.Errorf("Expected token secret to match, got mismatched token secret")
		}
	})
}
