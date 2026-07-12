package usecase

import "github.com/Leli2004/API_Go_biblioteca/internal/api/book"

type DeleteUC struct {
	repo book.Repository
}

func NewDeleteUC(repo book.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) error {
	return u.repo.Delete(id)
}
