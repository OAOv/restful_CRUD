package service

import (
	"github.com/OAOv/restful_CRUD/repository"
	"github.com/OAOv/restful_CRUD/types"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) CreateUser(user types.User) error {
	err := repository.CreateUser(user)
	return err
}

func (u *UserService) GetUsers() ([]types.User, error) {
	users, err := repository.GetUsers()
	return users, err
}

func (u *UserService) GetUser(id string) (types.User, error) {
	user, err := repository.GetUser(id)
	return user, err
}

func (u *UserService) UpdateUser(user types.User) error {
	err := repository.UpdateUser(user)
	return err
}

func (u *UserService) DeleteUser(id string) error {
	err := repository.DeleteUser(id)
	return err
}
