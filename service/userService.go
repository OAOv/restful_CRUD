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

func (s *Service) UpdateUser(user types.User) error {
	err := s.userRepository.UpdateUser(user)
	if user.Name != "" {
		err = s.recordRepository.UpdateReocrd(types.Record{UserID: user.ID, UserName: user.Name})
	}
	return err
}

func (s *Service) DeleteUser(id string) error {
	err := s.userRepository.DeleteUser(id)
	return err
}
