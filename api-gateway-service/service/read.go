package service

import (
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/api-gateway-service/entity"
	"github.com/rafimuhammad01/api-gateway-service/entity/dto"
	"github.com/rafimuhammad01/api-gateway-service/repository"
	"strconv"
	"time"
)

type ReadService interface {
	Read(request dto.ReadRequest) (*dto.TransactionReadResponse, error)
	ReadWithTransaction(request dto.ReadRequest) (*dto.TransactionReadResponse, error)
}

type readService struct {
	repository repository.ReadRepository
}

func NewReadService(repository repository.ReadRepository) ReadService {
	return &readService{repository: repository}
}

func (s *readService) ReadWithTransaction(request dto.ReadRequest) (*dto.TransactionReadResponse, error) {
	if request.TransactionID <= 0 {
		return nil, errors.Wrap(entity.ErrInvalidTrxID, "transaction id must be positive integer")
	}

	student, err := s.repository.ReadWithTransaction(request)
	if err != nil {
		return nil, err
	}

	if student.NPM != request.NPM {
		return nil, errors.Wrap(entity.ErrStudentNotFound, "student with npm "+request.NPM+" and transaction id "+strconv.Itoa(request.TransactionID)+" is not found in cache, please use /read endpoint")
	}

	return &dto.TransactionReadResponse{
		TransactionID: request.TransactionID,
		Student:       *student,
	}, nil
}

func (s *readService) Read(request dto.ReadRequest) (*dto.TransactionReadResponse, error) {
	trxID := int(time.Now().Unix())

	student, err := s.repository.Read(request, trxID)
	if err != nil {
		return nil, err
	}

	return &dto.TransactionReadResponse{
		TransactionID: trxID,
		Student: dto.ReadResponse{
			ID:   student.ID,
			NPM:  student.NPM,
			Name: student.Name,
		},
	}, nil
}
