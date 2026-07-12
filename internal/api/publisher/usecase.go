package publisher

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type UseCase interface {
	List(input entity.PublisherFilters) (error, entity.PublisherList)
	Get(id int) (error, entity.Publisher)
	Create(input entity.Publisher) (error, entity.Publisher)
	Update(id int, input entity.Publisher) (error, entity.Publisher)
	Delete(id int) error
}
