package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/reservation"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/Leli2004/API_Go_biblioteca/internal/helpers"
	"github.com/jmoiron/sqlx"
)

type ReservationUC struct {
	db   *sqlx.DB
	repo reservation.Repository
}

func NewUseCase(db *sqlx.DB, repo reservation.Repository) *ReservationUC {
	return &ReservationUC{db: db, repo: repo}
}

func (u *ReservationUC) Create(ctx context.Context, input entity.Reservation) (returnedCtx context.Context, err error, result entity.Reservation) {
	tx, err := helpers.OpenTransaction(ctx, u.db)
	if err != nil {
		return ctx, err, result
	}
	defer helpers.CloseTransaction(tx, &err)

	input.SetDefault()
	if err = input.Validate(); err != nil {
		return ctx, err, result
	}

	_, activeErr, active := u.repo.GetActiveByUserAndBook(ctx, tx, input.UserId, input.BookId)
	if activeErr == nil && active.Id != 0 {
		return ctx, errors.New("user already has an active reservation for this book"), result
	}

	if activeErr != nil && !errors.Is(activeErr, sql.ErrNoRows) {
		return ctx, activeErr, result
	}

	if input.Status == entity.ReservationStatusWaiting {
		_, err, position := u.repo.GetNextPosition(ctx, tx, input.BookId)
		if err != nil {
			return ctx, fmt.Errorf("ReservationUC.GetNextPosition: %w", err), result
		}
		input.Position = &position
	} else {
		input.Position = nil
	}

	returnedCtx, err, result = u.repo.Create(ctx, tx, input)
	if err != nil {
		return ctx, fmt.Errorf("ReservationUC.Create: %w", err), result
	}

	return
}
