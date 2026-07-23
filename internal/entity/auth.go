package entity

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginRequest) Validate() error {
	if l.Username == "" {
		return fmt.Errorf("invalid field: username is required")
	}

	if l.Password == "" {
		return fmt.Errorf("invalid field: password is required")
	}

	return nil
}

type AuthUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Token     string   `json:"token"`
	TokenType string   `json:"token_type"`
	ExpiresIn int64    `json:"expires_in"`
	User      AuthUser `json:"user"`
}

type AuthClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
