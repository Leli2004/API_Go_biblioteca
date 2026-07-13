package usecase

import (
	"errors"

	"github.com/Leli2004/API_Go_biblioteca/internal/api/loan"
	"github.com/Leli2004/API_Go_biblioteca/internal/entity"
)

type CreateUC struct {
	repo loan.Repository
}

func NewCreateUC(repo loan.Repository) CreateUC {
	return CreateUC{repo: repo}
}

func (u *CreateUC) Execute(input entity.Loan) (error, entity.Loan) {
	input.SetDefault()

	if err := input.Validate(); err != nil {
		return err, entity.Loan{}
	}

	errChk, existing := u.repo.GetActiveByUserAndBookCopy(input.UserId, input.BookCopyId)
	if errChk == nil && existing.Id != 0 {
		return errors.New("book_copy is already loaned"), entity.Loan{}
	}

	err, created := u.repo.Create(input)
	if err != nil {
		return err, entity.Loan{}
	}

	return nil, created
}
