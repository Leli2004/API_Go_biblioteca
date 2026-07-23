package security

import (
	"errors"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

var ErrUnauthorized = errors.New("unauthorized")

func ValidateRoles(claims *entity.AuthClaims, allowedRoles ...string) error {
	if claims == nil {
		return ErrUnauthorized
	}

	for _, allowedRole := range allowedRoles {
		if claims.Role == allowedRole {
			return nil
		}
	}

	return ErrUnauthorized
}
