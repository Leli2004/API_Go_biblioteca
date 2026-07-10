package author

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type UseCase interface {
	List(input entity.AuthorFilters) (error, entity.AuthorList)
}
