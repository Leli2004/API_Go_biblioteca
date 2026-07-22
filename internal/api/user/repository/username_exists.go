package repository

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UsernameExistsRepo struct{}

func NewUsernameExistsRepo() UsernameExistsRepo {
	return UsernameExistsRepo{}
}

func (r UsernameExistsRepo) Execute(ctx context.Context, tx *sqlx.Tx, username string, excludeUserID int) (context.Context, error, bool) {
	var exists bool

	err := tx.GetContext(ctx, &exists, usernameExistsSQL, strings.TrimSpace(username), excludeUserID)
	if err != nil {
		return ctx, err, false
	}

	return ctx, nil, exists
}

const usernameExistsSQL = `
	SELECT EXISTS (
		SELECT 1
		FROM biblioteca.users
		WHERE LOWER(username) = LOWER($1)
		  AND ($2 = 0 OR id <> $2)
	)
`
