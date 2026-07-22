package reservation

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

//go:generate mockery --name=UseCase --dir=. --output=mocks --filename=mock_usecase.go --with-expecter=True

type UseCase interface {
	Create(ctx context.Context, input entity.Reservation) (context.Context, error, entity.Reservation)
}
