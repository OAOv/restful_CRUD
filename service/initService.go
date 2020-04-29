package service

import (
	"github.com/OAOv/restful_CRUD/repository"
)

type Service struct {
	userRepository   repository.UserRepository
	recordRepository repository.RecordRepository
}
