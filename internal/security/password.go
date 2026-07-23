package security

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinimumPasswordLength = 4
	MaximumPasswordLength = 72
)

var (
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordTooShort = fmt.Errorf("password must contain at least %d characters", MinimumPasswordLength)
	ErrPasswordTooLong = fmt.Errorf("password must contain at most %d bytes", MaximumPasswordLength)
	ErrInvalidCredentials = errors.New("invalid username or password")
)

func ValidatePassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < MinimumPasswordLength {
		return ErrPasswordTooShort
	}

	// O bcrypt aceita senhas com no máximo 72 bytes
	if len([]byte(password)) > MaximumPasswordLength {
		return ErrPasswordTooLong
	}

	return nil
}

func HashPassword(password string) (string, error) {
	if err := ValidatePassword(password); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("security.HashPassword: %w", err)
	}

	return string(hash), nil
}

func ComparePassword(passwordHash, password string) error {
	if passwordHash == "" || password == "" {
		return ErrInvalidCredentials
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	if err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
