package genre

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type Repository interface {
	List(input entity.GenreFilters) (error, entity.GenreList)
}
