package usecase

import "github.com/Leli2004/API_Go_biblioteca/internal/api/user"

type DeleteUC struct {
	repo user.Repository
}

func NewDeleteUC(repo user.Repository) DeleteUC {
	return DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(id int) error {
	return u.repo.Delete(id)
}
