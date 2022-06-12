package service

import (
	"github.com/rafimuhammad01/read-service/entity/dto"
	"github.com/rafimuhammad01/read-service/repository"
)

type readService struct {
	repository repository.ReadRepository
}

type ReadService interface {
	Read(request dto.ReadRequest) (*dto.ReadResponse, error)
}

func NewReadService(repository repository.ReadRepository) ReadService {
	return &readService{
		repository: repository,
	}
}

func (s *readService) Read(npm dto.ReadRequest) (*dto.ReadResponse, error) {
	studentInfo, err := s.repository.Get(npm.NPM)
	if err != nil {
		return nil, err
	}

	return &dto.ReadResponse{
		ID:   studentInfo.ID,
		NPM:  studentInfo.NPM,
		Name: studentInfo.Name,
	}, nil
}
