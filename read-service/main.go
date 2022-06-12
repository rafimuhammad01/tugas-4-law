package main

import (
	"fmt"
	"github.com/rafimuhammad01/read-service/grpc/gen"
	"github.com/rafimuhammad01/read-service/handler"
	"github.com/rafimuhammad01/read-service/postgres"
	"github.com/rafimuhammad01/read-service/repository"
	"github.com/rafimuhammad01/read-service/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("READ_PORT")))
	if err != nil {
		logrus.Fatalf("Failed to connect to port %s", os.Getenv("READ_PORT"))
	}

	var opts []grpc.ServerOption

	db := postgres.Init()

	readRepo := repository.NewReadRepository(db)
	readService := service.NewReadService(readRepo)
	readHandler := handler.NewReadHandler(readService)

	grpcServer := grpc.NewServer(opts...)
	gen.RegisterReadServer(grpcServer, readHandler)

	logrus.Infof("Starting to run server on port %s", os.Getenv("READ_PORT"))
	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Fatalf("Failed to connect to port %s", os.Getenv("READ_PORT"))
	}
}
