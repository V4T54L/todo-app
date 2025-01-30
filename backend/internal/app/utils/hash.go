package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
	"todo_app_backend/internal/app/models"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed password.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CheckPasswordHash compares a hashed password with a plaintext password.
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateToken(tokenDetails models.Token) (string, error) {
	data, err := json.Marshal(tokenDetails)
	if err != nil {
		return "", err
	}

	// TODO: Add logic for token's authenticity

	encodedHash := base64.StdEncoding.EncodeToString(data)
	return encodedHash, nil
}

func ValidateToken(tokenStr string) (*models.Token, error) {

	// TODO: Add logic to verify token's authenticity

	var token models.Token
	err := json.Unmarshal([]byte(tokenStr), &token)
	if err != nil {
		return nil, fmt.Errorf("invalid token : %w", err)
	}

	if token.Exp.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return &token, nil
}
