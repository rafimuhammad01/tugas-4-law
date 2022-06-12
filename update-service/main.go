package main

import (
	"fmt"
	"github.com/rafimuhammad01/update-service/grpc/gen"
	"github.com/rafimuhammad01/update-service/handler"
	"github.com/rafimuhammad01/update-service/postgres"
	"github.com/rafimuhammad01/update-service/repository"
	"github.com/rafimuhammad01/update-service/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("UPDATE_PORT")))
	if err != nil {
		logrus.Fatalf("Failed to connect to port %s", os.Getenv("UPDATE_PORT"))
	}

	var opts []grpc.ServerOption

	db := postgres.Init()

	updateRepo := repository.NewUpdateRepository(db)
	updateService := service.NewUpdateService(updateRepo)
	updateHandler := handler.NewUpdateHandler(updateService)

	grpcServer := grpc.NewServer(opts...)
	gen.RegisterUpdateServer(grpcServer, updateHandler)

	logrus.Infof("Starting to run server on port %s", os.Getenv("UPDATE_PORT"))
	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Fatalf("Failed to connect to port %s", os.Getenv("UPDATE_PORT"))
	}
}
