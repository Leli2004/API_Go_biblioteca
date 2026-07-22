package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/labstack/echo"
)

const AuthClaimsContextKey = "auth_claims"

var ErrAuthClaimsNotFound = errors.New("authentication claims not found")

type JWTMiddleware struct {
	tokenManager *security.TokenManager
}

func NewJWTMiddleware(tokenManager *security.TokenManager) *JWTMiddleware {
	return &JWTMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *JWTMiddleware) Handler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if m == nil || m.tokenManager == nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "authentication middleware is not configured")
			}

			tokenValue, err := extractBearerToken(c.Request().Header.Get(echo.HeaderAuthorization))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			claims, err := m.tokenManager.Parse(tokenValue)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			c.Set(AuthClaimsContextKey, claims)

			return next(c)
		}
	}
}

func GetAuthClaims(c echo.Context) (*entity.AuthClaims, error) {
	if c == nil {
		return nil, ErrAuthClaimsNotFound
	}

	value := c.Get(AuthClaimsContextKey)

	claims, ok := value.(*entity.AuthClaims)
	if !ok || claims == nil {
		return nil, ErrAuthClaimsNotFound
	}

	return claims, nil
}

func RequireRoles(roles ...string) echo.MiddlewareFunc {
	allowedRoles := make(map[string]struct{}, len(roles))

	for _, role := range roles {
		role = strings.TrimSpace(role)

		if role != "" {
			allowedRoles[role] = struct{}{}
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := GetAuthClaims(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			if _, allowed := allowedRoles[claims.Role]; !allowed {
				return echo.NewHTTPError(http.StatusForbidden, "access denied")
			}

			return next(c)
		}
	}
}

func extractBearerToken(authorizationHeader string) (string, error) {
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 2 {
		return "", security.ErrInvalidToken
	}

	if !strings.EqualFold(fields[0], security.TokenTypeBearer) {
		return "", security.ErrInvalidToken
	}

	tokenValue := strings.TrimSpace(fields[1])
	if tokenValue == "" {
		return "", security.ErrInvalidToken
	}

	return tokenValue, nil
}
