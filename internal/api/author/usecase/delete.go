package usecase

import "github.com/Leli2004/API_Go_biblioteca/internal/api/author"

type DeleteUC struct {
	repo author.Repository
}

func NewDeleteUC(repo author.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) error {
	return u.repo.Delete(id)
}
