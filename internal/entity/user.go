package entity

import "fmt"

type User struct {
	Id           int     `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	Email        string  `json:"email" db:"email"`
	PasswordHash string  `json:"password_hash" db:"password_hash"`
	Phone        *string `json:"phone,omitempty" db:"phone"`
	Role         string  `json:"role,omitempty" db:"role"`
	Active       *bool   `json:"active,omitempty" db:"active"`
	CreatedAt    *string `json:"created_at" db:"created_at"`
	UpdatedAt    *string `json:"updated_at" db:"updated_at"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("Invalid field: Name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("Invalid field: Email is required")
	}
	if u.PasswordHash == "" {
		return fmt.Errorf("Invalid field: PasswordHash is required")
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
