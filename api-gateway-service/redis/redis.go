package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"os"
)

func Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Printf("[error to connect with Redis] %s", err.Error())
		return nil
	}

	logrus.Info("Redis Connected Successfully")
	return rdb
}
