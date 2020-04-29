package service

import "github.com/OAOv/restful_CRUD/types"

func (s *Service) CreateUser(user types.User) error {
	err := s.userRepository.CreateUser(user)
	return err
}

func (s *Service) GetUsers() ([]types.User, error) {
	users, err := s.userRepository.GetUsers()
	return users, err
}

func (s *Service) GetUser(id string) (types.User, error) {
	user, err := s.userRepository.GetUser(id)
	return user, err
}

func (s *Service) UpdateUser(id string, user map[string]interface{}) error {
	err := s.userRepository.UpdateUser(id, user)
	return err
}

func (s *Service) DeleteUser(id string) error {
	err := s.userRepository.DeleteUser(id)
	return err
}
