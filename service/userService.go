package service

import (
	"github.com/OAOv/restful_CRUD/repository"
	"github.com/OAOv/restful_CRUD/types"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) GetUsers() ([]types.User, error) {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
