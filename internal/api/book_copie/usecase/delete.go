package usecase

import book_copie "github.com/Leli2004/API_Go_biblioteca/internal/api/book_copie"

type DeleteUC struct {
	repo book_copie.Repository
}

func NewDeleteUC(repo book_copie.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) error {
	return u.repo.Delete(id)
}
