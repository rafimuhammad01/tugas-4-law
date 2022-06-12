package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rafimuhammad01/api-gateway-service/grpc/gen"
	"github.com/rafimuhammad01/api-gateway-service/handler"
	"github.com/rafimuhammad01/api-gateway-service/redis"
	"github.com/rafimuhammad01/api-gateway-service/repository"
	"github.com/rafimuhammad01/api-gateway-service/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func main() {
	e := echo.New()

	// external db
	cache := redis.Init()

	// grpc package
	updateConn, err := grpc.Dial(fmt.Sprintf("dns:///%s:%s", os.Getenv("UPDATE_HOST"), os.Getenv("UPDATE_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("[grpc error] - %s", err)
	}
	updateClient := gen.NewUpdateClient(updateConn)

	readConn, err := grpc.Dial(fmt.Sprintf("dns:///%s:%s", os.Getenv("READ_HOST"), os.Getenv("READ_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	readClient := gen.NewReadClient(readConn)

	// internal package
	updateRepository := repository.NewUpdateRepository(cache, updateClient)
	updateService := service.NewUpdateService(updateRepository)
	updateHandler := handler.NewUpdateHandler(updateService)

	readRepository := repository.NewReadRepository(cache, readClient)
	readService := service.NewReadService(readRepository)
	readHandler := handler.NewReadHandler(readService)

	e.POST("/update", updateHandler.Create)
	e.GET("/read/:npm", readHandler.Read)
	e.GET("/read/:npm/:trxID", readHandler.ReadWithTransaction)

	logrus.Infof("Starting to run server on port %s", os.Getenv("GATEWAY_PORT"))
	err = e.Start(fmt.Sprintf(":%s", os.Getenv("GATEWAY_PORT")))
	if err != nil {
		logrus.Fatalf("[grpc error] - %s", err)
	}
}
