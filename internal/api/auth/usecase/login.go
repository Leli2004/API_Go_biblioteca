package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	userAPI "github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
)

type LoginUC struct {
	db           *sqlx.DB
	userRepo     userAPI.Repository
	tokenManager *security.TokenManager
}

func NewLoginUC(db *sqlx.DB, userRepo userAPI.Repository, tokenManager *security.TokenManager) LoginUC {
	return LoginUC{
		db:           db,
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (u LoginUC) Execute(ctx context.Context, input entity.LoginRequest) (returnedCtx context.Context, err error, result entity.LoginResponse) {
	input.Username = strings.TrimSpace(input.Username)

	if err := input.Validate(); err != nil {
		return ctx, fmt.Errorf("AuthUC.Login.Validate: %w", err), entity.LoginResponse{}
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, fmt.Errorf("AuthUC.Login.OpenTransaction: %w", err), entity.LoginResponse{}
	}

	defer helpers.CloseTransaction(tx, &err)

	returnedCtx, err, foundUser := u.userRepo.GetByUsername(ctx, tx, input.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return returnedCtx, security.ErrInvalidCredentials, entity.LoginResponse{}
		}

		return returnedCtx, fmt.Errorf("AuthUC.Login.GetByUsername: %w", err), entity.LoginResponse{}
	}

	if foundUser.Active == nil || !*foundUser.Active {
		return returnedCtx, security.ErrInvalidCredentials, entity.LoginResponse{}
	}

	if err = security.ComparePassword(foundUser.PasswordHash, input.Password); err != nil {
		return returnedCtx, security.ErrInvalidCredentials, entity.LoginResponse{}
	}

	tokenValue, _, err := u.tokenManager.Generate(foundUser.Id, foundUser.Username, foundUser.Role)
	if err != nil {
		return returnedCtx, fmt.Errorf("AuthUC.Login.GenerateToken: %w", err), entity.LoginResponse{}
	}

	result = entity.LoginResponse{
		Token:     tokenValue,
		TokenType: security.TokenTypeBearer,
		ExpiresIn: u.tokenManager.ExpirationSeconds(),
		User: entity.AuthUser{
			Id:       foundUser.Id,
			Name:     foundUser.Name,
			Username: foundUser.Username,
			Email:    foundUser.Email,
			Role:     foundUser.Role,
		},
	}

	return returnedCtx, nil, result
}
