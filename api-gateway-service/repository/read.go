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
	"strconv"
	"strings"
	"time"
)

type readRepository struct {
	cache      *redis.Client
	readClient gen.ReadClient
}

func NewReadRepository(cache *redis.Client, grpcReadClient gen.ReadClient) ReadRepository {
	return &readRepository{
		cache:      cache,
		readClient: grpcReadClient,
	}
}

type ReadRepository interface {
	Read(request dto.ReadRequest, trxID int) (*dto.ReadResponse, error)
	ReadWithTransaction(request dto.ReadRequest) (*dto.ReadResponse, error)
}

func (r *readRepository) ReadWithTransaction(request dto.ReadRequest) (*dto.ReadResponse, error) {
	ctx := context.Background()

	var studentFromCache entity.Student
	err := r.cache.Get(ctx, strconv.Itoa(request.TransactionID)).Scan(&studentFromCache)
	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(entity.ErrStudentNotFound, "student with npm "+request.NPM+" and transaction id "+strconv.Itoa(request.TransactionID)+" is not found in cache, please use /read endpoint")
		}
		return nil, err
	}

	return &dto.ReadResponse{
		ID:   studentFromCache.ID,
		NPM:  studentFromCache.NPM,
		Name: studentFromCache.Name,
	}, nil
}

func (r *readRepository) Read(request dto.ReadRequest, trxID int) (*dto.ReadResponse, error) {
	ctx := context.Background()
	resp, err := r.readClient.Read(ctx, &gen.ReadRequest{Npm: request.NPM})
	if err != nil {
		errStatus, _ := status.FromError(err)
		if errStatus.Code() == codes.NotFound {
			if strings.Contains(errStatus.Message(), entity.ErrStudentNotFound.Error()) {
				return nil, errors.Wrap(entity.ErrStudentNotFound, errStatus.Message())
			}
		}

		return nil, err
	}

	if err = r.cache.Set(ctx, strconv.Itoa(trxID), entity.Student{
		ID:   int(resp.Id),
		NPM:  resp.Npm,
		Name: resp.Name,
	}, 30*time.Minute).Err(); err != nil {
		return nil, err
	}

	return &dto.ReadResponse{
		ID:   int(resp.Id),
		NPM:  resp.Npm,
		Name: resp.Name,
	}, nil
}
