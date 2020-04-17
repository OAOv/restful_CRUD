package service

import "github.com/OAOv/restful_CRUD/types"

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) GetUsers() ([]types.User, error) {
	return nil, nil
}
