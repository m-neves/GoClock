package service

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/repository"
)

type UserServiceInterface interface {
	FindById(id int) (*data.User, error)
	FindByCredentials(email, password string) (*data.User, error)
	Create(user *data.User) error
}

type userService struct {
	r repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *userService {
	return &userService{repository}
}

func (u *userService) FindById(id int) (*data.User, error) {
	return u.r.FindById(id)
}

func (u *userService) FindByCredentials(email, password string) (*data.User, error) {
	password = SHA256Encoder(password)
	return u.r.FindByCredentials(email, password)
}

func (u *userService) Create(user *data.User) error {
	user.Password = SHA256Encoder(user.Password)

	return u.r.Create(user)
}
