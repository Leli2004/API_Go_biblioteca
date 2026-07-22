package repository

import (
	"context"
	"github.com/Leli2004/API_Go_biblioteca/internal/api/fine"
	"github.com/jmoiron/sqlx"
)

type FineRepo struct{}

func NewRepository() *FineRepo { return &FineRepo{} }

var _ fine.Repository = (*FineRepo)(nil)

func (r *FineRepo) ProcessOverdueLoans(ctx context.Context, tx *sqlx.Tx, amount float64, reason string) (context.Context, error, int) {
	var count int
	err := tx.GetContext(ctx, &count, processSQL, amount, reason)
	return ctx, err, count
}

var processSQL = `
WITH candidates AS (
 SELECT id,user_id FROM biblioteca.loans
 WHERE status='active' AND returned_at IS NULL AND due_date < NOW() AND user_id IS NOT NULL
 FOR UPDATE SKIP LOCKED
), updated AS (
 UPDATE biblioteca.loans l SET status='overdue',updated_at=NOW()
 FROM candidates c WHERE l.id=c.id
 RETURNING l.id,l.user_id
), inserted AS (
 INSERT INTO biblioteca.fines(loan_id,user_id,amount,reason)
 SELECT id,user_id,$1,$2 FROM updated
 ON CONFLICT (loan_id) DO NOTHING
 RETURNING id
)
SELECT COUNT(*) FROM inserted;`
