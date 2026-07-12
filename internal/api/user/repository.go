package user

import "github.com/Leli2004/API_Go_biblioteca/internal/entity"

type Repository interface {
	List(input entity.UserFilters) (error, entity.UserList)
	Get(id int) (error, entity.User)
	Create(input entity.User) (error, entity.User)
	Update(id int, input entity.User) (error, entity.User)
	Delete(id int) error
}
