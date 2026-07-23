package repository

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
)

type EmailExistsRepo struct{}

func NewEmailExistsRepo() EmailExistsRepo {
	return EmailExistsRepo{}
}

func (r EmailExistsRepo) Execute(ctx context.Context, tx *sqlx.Tx, email string, excludeUserID int) (context.Context, error, bool) {
	var exists bool

	err := tx.GetContext(ctx, &exists, emailExistsSQL, strings.TrimSpace(email), excludeUserID)
	if err != nil {
		return ctx, err, false
	}

	return ctx, nil, exists
}

const emailExistsSQL = `
	SELECT EXISTS (
		SELECT 1
		FROM biblioteca.users
		WHERE LOWER(email) = LOWER($1)
		  AND ($2 = 0 OR id <> $2)
	)
`
