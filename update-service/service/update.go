package service

import (
	"github.com/rafimuhammad01/update-service/entity"
	"github.com/rafimuhammad01/update-service/entity/dto"
	"github.com/rafimuhammad01/update-service/repository"
)

type updateService struct {
	repository repository.UpdateRepository
}

type UpdateService interface {
	Create(request dto.CreateStudentRequest) error
}

func NewUpdateService(repository repository.UpdateRepository) UpdateService {
	return &updateService{
		repository: repository,
	}
}

func (s *updateService) Create(request dto.CreateStudentRequest) error {
	student := entity.Student{
		NPM:  request.NPM,
		Name: request.Name,
	}

	if err := student.Validate(); err != nil {
		return err
	}

	if err := s.repository.Create(student); err != nil {
		return err
	}

	return nil
}
