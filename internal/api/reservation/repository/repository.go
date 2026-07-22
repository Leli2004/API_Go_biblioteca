package repository

import (
	"context"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/reservation"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ReservationRepo struct{}

func NewRepository() *ReservationRepo { return &ReservationRepo{} }

var _ reservation.Repository = (*ReservationRepo)(nil)

func (r *ReservationRepo) GetActiveByUserAndBook(ctx context.Context, tx *sqlx.Tx, userId, bookId int) (context.Context, error, entity.Reservation) {
	var result entity.Reservation
	err := tx.GetContext(ctx, &result, sqlGetActiveByUser, userId, bookId)
	return ctx, err, result
}

func (r *ReservationRepo) GetNextPosition(ctx context.Context, tx *sqlx.Tx, bookId int) (context.Context, error, int) {
	if _, err := tx.ExecContext(ctx, sqlGetNextPosition1, bookId); err != nil {
		return ctx, err, 0
	}
	var position int
	err := tx.GetContext(ctx, &position, sqlGetNextPosition2, bookId)
	return ctx, err, position
}

func (r *ReservationRepo) Create(ctx context.Context, tx *sqlx.Tx, input entity.Reservation) (context.Context, error, entity.Reservation) {
	var result entity.Reservation
	err := tx.GetContext(ctx, &result, sqlCreate, input.UserId, input.BookId, input.Position, input.Status, input.ExpiresAt)
	return ctx, err, result
}

var (
	sqlGetActiveByUser = `SELECT id,user_id,book_id,position,status,expires_at,created_at,updated_at FROM biblioteca.reservations WHERE user_id=$1 AND book_id=$2 AND status IN ('waiting','available') ORDER BY id DESC LIMIT 1`

	sqlGetNextPosition1 = `SELECT pg_advisory_xact_lock($1)`
	sqlGetNextPosition2 = `SELECT COALESCE(MAX(position),0)+1 FROM biblioteca.reservations WHERE book_id=$1 AND status='waiting'`

	sqlCreate = `INSERT INTO biblioteca.reservations(user_id,book_id,position,status,expires_at) VALUES($1,$2,$3,$4,$5::timestamptz) RETURNING id,user_id,book_id,position,status,expires_at,created_at,updated_at`
)
