package auth

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type UseCase interface {
	Login(ctx context.Context,input entity.LoginRequest) (context.Context,error,entity.LoginResponse)
	Me(ctx context.Context,userID int) (context.Context,error,entity.AuthUser)
}
