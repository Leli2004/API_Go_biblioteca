package usecase

import (
	"context"
	"fmt"

	userAPI "github.com/Leli2004/API_Go_biblioteca/internal/api/user"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type MeUC struct {
	db       *sqlx.DB
	userRepo userAPI.Repository
}

func NewMeUC(db *sqlx.DB, userRepo userAPI.Repository) MeUC {
	return MeUC{
		db:       db,
		userRepo: userRepo,
	}
}

func (u MeUC) Execute(ctx context.Context, userID int) (returnedCtx context.Context, err error, result entity.AuthUser) {
	if userID <= 0 {
		return ctx, fmt.Errorf("AuthUC.Me: invalid user id"), entity.AuthUser{}
	}

	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, fmt.Errorf("AuthUC.Me.OpenTransaction: %w", err), entity.AuthUser{}
	}

	defer helpers.CloseTransaction(tx, &err)

	returnedCtx, err, foundUser := u.userRepo.Get(ctx, tx, userID)
	if err != nil {
		return returnedCtx, fmt.Errorf("AuthUC.Me.GetUser: %w", err), entity.AuthUser{}
	}

	if foundUser.Active == nil || !*foundUser.Active {
		return returnedCtx, fmt.Errorf("AuthUC.Me: user is inactive"), entity.AuthUser{}
	}

	result = entity.AuthUser{
		Id:       foundUser.Id,
		Name:     foundUser.Name,
		Username: foundUser.Username,
		Email:    foundUser.Email,
		Role:     foundUser.Role,
	}

	return returnedCtx, nil, result
}
