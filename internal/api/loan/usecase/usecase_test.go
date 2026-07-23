package usecase

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mm "github.com/Leli2004/API_Go_biblioteca/internal/api/loan/mocks"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupLoan(t *testing.T) (*LoanUC, *mm.Repository, sqlmock.Sqlmock, context.Context) {
	s, m, e := sqlmock.New()
	assert.NoError(t, e)
	db := sqlx.NewDb(s, "sqlmock")
	r := mm.NewRepository(t)
	t.Cleanup(func() { _ = db.Close() })
	return NewUseCase(db, r), r, m, context.Background()
}

func Test_Get_UseCase(t *testing.T) {
	u, r, m, c := setupLoan(t)
	m.ExpectBegin()
	r.On("Get", mock.Anything, mock.Anything, 1).Return(c, nil, entity.Loan{}).Once()
	m.ExpectCommit()
	_, e, _ := u.Get(c, 1)
	assert.NoError(t, e)
}

func Test_List_UseCase(t *testing.T) {
	u, r, m, c := setupLoan(t)
	in := entity.LoanFilters{}
	x := in
	x.SetDefault()
	m.ExpectBegin()
	r.On("List", mock.Anything, mock.Anything, x).Return(c, nil, entity.LoanList{}).Once()
	m.ExpectCommit()
	_, e, _ := u.List(c, in)
	assert.NoError(t, e)
}

func Test_Return_UseCase(t *testing.T) {
	u, r, m, c := setupLoan(t)
	ln := entity.Loan{}
	m.ExpectBegin()
	r.On("Get", mock.Anything, mock.Anything, 1).Return(c, nil, ln).Once()
	r.On("Update", mock.Anything, mock.Anything, 1, mock.MatchedBy(func(v entity.Loan) bool { return v.Status == "returned" && v.ReturnedAt != nil })).Return(c, nil, ln).Once()
	m.ExpectCommit()
	claims := &entity.AuthClaims{Role: entity.RoleAdmin}
	_, e, _ := u.Return(c, 1, nil, claims)
	assert.NoError(t, e)
}
