package service

import (
	"github.com/OAOv/restful_CRUD/repository"
	"github.com/OAOv/restful_CRUD/types"
)

type Service struct {
	userRepository   repository.UserRepository
	recordRepository repository.RecordRepository
}

//// user part
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
	if err != nil {
		return nil
	}
	err = s.recordRepository.DeleteRecord(id, true)
	return err
}

//// record part
func (s *Service) CreateRecord(record types.Record) error {
	err := s.recordRepository.CreateRecord(record)
	return err
}

func (s *Service) GetRecords() ([]types.Record, error) {
	records, err := s.recordRepository.GetRecords()
	return records, err
}

func (s *Service) GetRecord(id string) (types.Record, error) {
	record, err := s.recordRepository.GetRecord(id)
	return record, err
}

func (s *Service) GetRecordByUser(id string) ([]types.Record, error) {
	records, err := s.recordRepository.GetRecordByUser(id)
	return records, err
}

func (s *Service) UpdateRecord(record types.Record) error {
	err := s.recordRepository.UpdateReocrd(record)
	return err
}

func (s *Service) DeleteRecord(id string, isUser bool) error {
	err := s.recordRepository.DeleteRecord(id, false)
	return err
}
