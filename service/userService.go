package service

import (
	"github.com/OAOv/restful_CRUD/repository"
	"github.com/OAOv/restful_CRUD/types"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) GetUsers() ([]types.User, error) {
	users, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) GetUser(id string) (types.User, error) {
	user, err := repository.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
