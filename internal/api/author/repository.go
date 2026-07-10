package author

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type Repository interface {
	List(input entity.AuthorFilters) (error, entity.AuthorList)
	Get(id int) (error, entity.Author)
	Create(input entity.Author) (error, entity.Author)
	Update(id int, input entity.Author) (error, entity.Author)
	Delete(id int) error
}
