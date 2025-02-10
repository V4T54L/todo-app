package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"todo_app_backend/config"
	"todo_app_backend/internal/app/models"

	"golang.org/x/crypto/bcrypt"
)

const (
	saperator byte = '^'
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateToken(tokenDetails models.Token) (string, error) {
	data, err := json.Marshal(tokenDetails)
	if err != nil {
		return "", fmt.Errorf("error marshaling token details: %w", err)
	}

	secret, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("error getting config: %w", err)
	}

	// Append the secret to data
	secretData := append(data, secret.TokenSecret...)

	// Calculate hash
	hashValue := sha256.Sum256(secretData)

	// Combine the raw data and the hash for the token
	data = append(data, saperator)
	data = append(data, hashValue[:]...)

	// Encode the final token as base64
	encodedHash := base64.URLEncoding.EncodeToString(data)
	return encodedHash, nil
}

func ValidateToken(tokenStr string) (*models.Token, error) {
	data, err := base64.URLEncoding.DecodeString(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding token: %w", err)
	}

	// Split at separator character
	idx := bytes.IndexByte(data, saperator)
	if idx == -1 {
		return nil, errors.New("invalid token format")
	}

	// Raw data and hash
	raw := data[:idx]
	hashed := data[idx+1:]

	// Get the secret for validation
	secret, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting config: %w", err)
	}

	// Validate authenticity
	expectedHash := sha256.Sum256(append(raw, secret.TokenSecret...))
	if !bytes.Equal(expectedHash[10:], hashed[10:]) {
		return nil, errors.New("invalid token hash")
	}

	var token models.Token
	if err = json.Unmarshal(raw, &token); err != nil {
		return nil, fmt.Errorf("error unmarshaling token: %w", err)
	}

	if token.Exp.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return &token, nil
}
