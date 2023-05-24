package repository

import (
	// "database/sql"

	"errors"
	"fmt"

	"github.com/lisyaoran51/GoSimpleApi/entity"
	"github.com/lisyaoran51/GoSimpleApi/utility"
)

type MemoryUserRepository struct {
	Users utility.GenericMap[string, entity.User]
}

func NewMemoryUserRepository() UserRepository {
	return &MemoryUserRepository{}
}

func (r *MemoryUserRepository) Save(user entity.User) error {
	if value, err := r.FindByAccount(user.Account); err == nil {
		fmt.Printf("LOG: user exist: %#v", value)
		return errors.New("user already exist")
	}
	r.Users.Store(user.Account, user)
	return nil
}

func (r *MemoryUserRepository) Update(user entity.User) error {
	r.Users.Store(user.Account, user)
	return nil
}

func (r *MemoryUserRepository) Delete(user entity.User) error {
	r.Users.Delete(user.Account)

	return nil
}

func (r *MemoryUserRepository) FindAll() []entity.User {

	return nil
}

func (r *MemoryUserRepository) FindByID(id int) (entity.User, error) {

	return entity.User{}, nil
}

func (r *MemoryUserRepository) FindByAccount(account string) (entity.User, error) {
	value, ok := r.Users.Load(account)
	if !ok {
		return value, errors.New("not found")
	}
	return value, nil
}
