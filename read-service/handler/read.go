package handler

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/read-service/entity"
	"github.com/rafimuhammad01/read-service/entity/dto"
	"github.com/rafimuhammad01/read-service/grpc/gen"
	"github.com/rafimuhammad01/read-service/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReadHandler struct {
	gen.UnimplementedReadServer
	service service.ReadService
}

func NewReadHandler(service service.ReadService) *ReadHandler {
	return &ReadHandler{service: service}
}

func (h *ReadHandler) Read(ctx context.Context, in *gen.ReadRequest) (*gen.ReadResponse, error) {
	resp, err := h.service.Read(dto.ReadRequest{NPM: in.Npm})
	if err != nil {
		if errors.Cause(err) == entity.ErrStudentNotFound {
			return nil, status.Errorf(
				codes.NotFound,
				err.Error(),
			)
		}

		logrus.Error(err.Error())
		return nil, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	return &gen.ReadResponse{Id: int32(resp.ID), Npm: resp.NPM, Name: resp.Name}, nil
}
