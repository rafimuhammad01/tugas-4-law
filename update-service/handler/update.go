package handler

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/update-service/entity"
	"github.com/rafimuhammad01/update-service/entity/dto"
	"github.com/rafimuhammad01/update-service/grpc/gen"
	"github.com/rafimuhammad01/update-service/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateHandler struct {
	gen.UnimplementedUpdateServer
	service service.UpdateService
}

func NewUpdateHandler(service service.UpdateService) *UpdateHandler {
	return &UpdateHandler{service: service}
}

func (h *UpdateHandler) Create(ctx context.Context, in *gen.CreateRequest) (*gen.CreateResponse, error) {
	err := h.service.Create(dto.CreateStudentRequest{
		NPM:  in.Npm,
		Name: in.Name,
	})
	if err != nil {
		if errors.Cause(err) == entity.ErrInvalidName || errors.Cause(err) == entity.ErrInvalidNPM {
			return nil, status.Error(
				codes.InvalidArgument,
				err.Error(),
			)
		}
		logrus.Error(err.Error())
		return nil, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	return &gen.CreateResponse{}, nil
}
