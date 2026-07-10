package genre

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type Repository interface {
	List(input entity.GenreFilters) (error, entity.GenreList)
	Get(id int) (error, entity.Genre)
	Create(input entity.Genre) (error, entity.Genre)
	Update(id int, input entity.Genre) (error, entity.Genre)
	Delete(id int) error
}
