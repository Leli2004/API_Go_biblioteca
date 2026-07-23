package security

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

const (
	DefaultTokenIssuer     = "biblioteca-api"
	DefaultTokenExpiration = 24 * time.Hour
	TokenTypeBearer        = "Bearer"
)

var (
	ErrJWTSecretRequired = errors.New("JWT secret is required")
	ErrInvalidToken      = errors.New("invalid or expired token")
)

type TokenManager struct {
	secret     []byte
	issuer     string
	expiration time.Duration
	now        func() time.Time
}

func NewTokenManager(
	secret string,
	issuer string,
	expiration time.Duration,
) (*TokenManager, error) {
	if secret == "" {
		return nil, ErrJWTSecretRequired
	}

	if issuer == "" {
		issuer = DefaultTokenIssuer
	}

	if expiration <= 0 {
		expiration = DefaultTokenExpiration
	}

	return &TokenManager{
		secret:     []byte(secret),
		issuer:     issuer,
		expiration: expiration,
		now:        time.Now,
	}, nil
}

func (m *TokenManager) Generate(userId int, username string, role string) (tokenValue string, expiresAt time.Time, err error) {
	if userId <= 0 {
		return "", time.Time{}, fmt.Errorf("security.TokenManager.Generate: invalid user id")
	}

	if username == "" {
		return "", time.Time{}, fmt.Errorf("security.TokenManager.Generate: username is required")
	}

	if role != entity.RoleAdmin && role != entity.RoleUser {
		return "", time.Time{}, fmt.Errorf("security.TokenManager.Generate: invalid user role")
	}

	now := m.now().UTC()
	expiresAt = now.Add(m.expiration)

	claims := entity.AuthClaims{
		UserId:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(userId),
			Issuer:    m.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err = token.SignedString(m.secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("security.TokenManager.Generate: %w", err)
	}

	return tokenValue, expiresAt, nil
}

func (m *TokenManager) Parse(tokenValue string) (*entity.AuthClaims, error) {
	if tokenValue == "" {
		return nil, ErrInvalidToken
	}

	claims := &entity.AuthClaims{}

	token, err := jwt.ParseWithClaims(
		tokenValue,
		claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, ErrInvalidToken
			}

			return m.secret, nil
		},
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
		jwt.WithIssuer(m.issuer),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
	)

	if err != nil || token == nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims.UserId <= 0 ||
		claims.Username == "" ||
		(claims.Role != entity.RoleAdmin &&
			claims.Role != entity.RoleUser) {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func (m *TokenManager) ExpirationSeconds() int64 {
	return int64(m.expiration.Seconds())
}
