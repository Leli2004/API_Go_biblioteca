package book

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type UseCase interface {
	List(input entity.BookFilters) (error, entity.BookList)
	Get(id int) (error, entity.Book)
	Create(input entity.Book) (error, entity.Book)
	Update(id int, input entity.Book) (error, entity.Book)
	Delete(id int) error
}
