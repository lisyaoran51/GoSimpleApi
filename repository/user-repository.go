package repository

import (
	"github.com/lisyaoran51/GoSimpleApi/entity"
)

type UserRepository interface {
	Save(user entity.User) error
	Update(user entity.User) error
	Delete(user entity.User) error
	FindAll() []entity.User
	FindByID(id int) (entity.User, error)
	FindByAccount(account string) (entity.User, error)
}
