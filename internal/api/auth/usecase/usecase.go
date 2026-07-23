package usecase

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/auth"
	userAPI "github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/security"
	"github.com/jmoiron/sqlx"
)

type UseCase struct {
	db           *sqlx.DB
	userRepo     userAPI.Repository
	tokenManager *security.TokenManager
}

func NewUseCase(db *sqlx.DB, userRepo userAPI.Repository, tokenManager *security.TokenManager) auth.UseCase {
	return &UseCase{
		db:           db,
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (u *UseCase) Login(ctx context.Context, input entity.LoginRequest) (context.Context, error, entity.LoginResponse) {
	return NewLoginUC(u.db, u.userRepo, u.tokenManager).Execute(ctx, input)
}

func (u *UseCase) Me(ctx context.Context, userID int) (context.Context, error, entity.AuthUser) {
	return NewMeUC(u.db, u.userRepo).Execute(ctx, userID)
}
