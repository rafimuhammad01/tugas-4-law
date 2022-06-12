package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/api-gateway-service/entity"
	"github.com/rafimuhammad01/api-gateway-service/entity/dto"
	"github.com/rafimuhammad01/api-gateway-service/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type updateRepository struct {
	cache        *redis.Client
	updateClient gen.UpdateClient
}

func NewUpdateRepository(cache *redis.Client, grpcUpdateClient gen.UpdateClient) UpdateRepository {
	return &updateRepository{
		cache:        cache,
		updateClient: grpcUpdateClient,
	}
}

type UpdateRepository interface {
	Create(request dto.CreateRequest) error
}

func (r *updateRepository) Create(request dto.CreateRequest) error {
	_, err := r.updateClient.Create(context.Background(), &gen.CreateRequest{Name: request.Name, Npm: request.NPM})
	if err != nil {
		errStatus, _ := status.FromError(err)
		if errStatus.Code() == codes.InvalidArgument {
			if strings.Contains(errStatus.Message(), entity.ErrInvalidNPM.Error()) {
				return errors.Wrap(entity.ErrInvalidNPM, errStatus.Message())
			}

			if strings.Contains(errStatus.Message(), entity.ErrInvalidName.Error()) {
				return errors.Wrap(entity.ErrInvalidName, errStatus.Message())
			}
		}

		return err
	}

	return nil
}
