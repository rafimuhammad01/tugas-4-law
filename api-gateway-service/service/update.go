package service

import (
	"github.com/rafimuhammad01/api-gateway-service/entity/dto"
	"github.com/rafimuhammad01/api-gateway-service/repository"
)

type UpdateService interface {
	Create(request dto.CreateRequest) error
}

type updateService struct {
	repository repository.UpdateRepository
}

func NewUpdateService(repository repository.UpdateRepository) UpdateService {
	return &updateService{repository: repository}
}

func (s *updateService) Create(request dto.CreateRequest) error {
	if err := s.repository.Create(request); err != nil {
		return err
	}

	return nil
}
