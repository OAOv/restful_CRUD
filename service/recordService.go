package service

import "github.com/OAOv/restful_CRUD/types"

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

func (s *Service) UpdateRecord(id string, record map[string]interface{}) error {
	err := s.recordRepository.UpdateReocrd(id, record)
	return err
}

func (s *Service) DeleteRecord(id string) error {
	err := s.recordRepository.DeleteRecord(id)
	return err
}
