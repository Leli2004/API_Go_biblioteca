package book_copie

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type Repository interface {
	List(input entity.BookCopyFilters) (error, entity.BookCopyList)
	Get(id int) (error, entity.BookCopy)
	Create(input entity.BookCopy) (error, entity.BookCopy)
	Update(id int, input entity.BookCopy) (error, entity.BookCopy)
	Delete(id int) error
}
