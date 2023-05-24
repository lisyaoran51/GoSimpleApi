package service

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/lisyaoran51/GoSimpleApi/entity"
	"github.com/lisyaoran51/GoSimpleApi/repository"
	"github.com/pkg/errors"
)

type UserService interface {
	Save(entity.User) (entity.User, error)
	Validate(entity.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) Save(user entity.User) (entity.User, error) {

	fmt.Printf("LOG: username: %s\n", user.Account)
	match, err := regexp.MatchString("[0-9A-Za-z]", user.Account)
	if err != nil {
		return entity.User{}, errors.WithStack(err)
	}
	if !match {
		return entity.User{}, errors.New("unsupported characters in account. only alphanum supported.")
	}
	if len(user.Account) < 3 || len(user.Account) > 32 {
		return entity.User{}, errors.New("unsupported length in account.")
	}

	fmt.Printf("LOG: password: %s\n", user.Password)
	match, err = regexp.MatchString("[0-9A-Za-z]", user.Password)
	if err != nil {
		return entity.User{}, errors.WithStack(err)
	}
	if !match {
		return entity.User{}, errors.New("unsupported characters in password. only alphanum supported.")
	}
	if len(user.Password) < 8 || len(user.Password) > 32 {
		return entity.User{}, errors.New("unsupported length in password.")
	}

	hasUpperCase := false
	hasLowerCase := false
	hasNumber := false
	for _, r := range user.Password {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			hasUpperCase = true
		}
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			hasLowerCase = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
	}

	if !hasLowerCase || !hasUpperCase || !hasNumber {
		return entity.User{}, errors.New("need 1 uppercase, 1 lowercase, 1 number.")
	}

	return user, service.userRepository.Save(user)
}

func (service *userService) Validate(user entity.User) error {

	storedUser, err := service.userRepository.FindByAccount(user.Account)
	if err != nil {
		if err.Error() == "not found" {
			return errors.New("user not exist")
		}
		return err
	}

	if storedUser.Password != user.Password {
		return errors.New("password not same")
	}

	return nil
}
