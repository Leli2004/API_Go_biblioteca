package helpers

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func OpenTransaction(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, error) {
	return db.BeginTxx(ctx, nil)
}

func CloseTransaction(tx *sqlx.Tx, errRef *error) {
	if recovered := recover(); recovered != nil {
		_ = tx.Rollback()
		panic(recovered)
	}

	if *errRef != nil {
		_ = tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		*errRef = err
	}
}
