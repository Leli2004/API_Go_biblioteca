package entity

import (
	"fmt"
	"strings"
)

type User struct {
	Id           int     `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	Email        string  `json:"email" db:"email"`
	Username     string  `json:"username" db:"username"`
	Password     string  `json:"password,omitempty" db:"-"`
	PasswordHash string  `json:"password_hash" db:"password_hash"`
	Phone        *string `json:"phone,omitempty" db:"phone"`
	Role         string  `json:"role,omitempty" db:"role"`
	Active       *bool   `json:"active,omitempty" db:"active"`
	CreatedAt    *string `json:"created_at" db:"created_at"`
	UpdatedAt    *string `json:"updated_at" db:"updated_at"`
}

func (u *User) SetDefault() {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	if u.Role == "" {
		u.Role = RoleUser
	}

	if u.Active == nil {
		active := true
		u.Active = &active
	}
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("Invalid field: Name is required")
	}
	if u.Username == "" {
		return fmt.Errorf("invalid field: username is required")
	}
	if strings.ContainsAny(u.Username, " \t\n") {
		return fmt.Errorf("invalid field: username must not contain spaces")
	}
	if u.Email == "" {
		return fmt.Errorf("Invalid field: Email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("Invalid field: Password is required")
	}
	if u.Role != RoleAdmin && u.Role != RoleUser {
		return fmt.Errorf("invalid field: role must be admin or user")
	}
	return nil
}

type UserFilters struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (u *UserFilters) SetDefault() {
	if u.Limit == 0 {
		u.Limit = 10
	}
}

type UserList struct {
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Data   []*User `json:"data"`
}
