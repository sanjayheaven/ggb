package redis

import (
	"fmt"

	"github.com/sanjayheaven/ggb/configs"
	"github.com/sanjayheaven/ggb/internal/pkg/logger"

	"github.com/redis/go-redis/v9"
)

func Connect(config *configs.Redis) *redis.Client {

	logger := logger.LogrusLogger

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.UserName,
		Password: config.Password,
		DB:       config.Db,
	})

	logger.Printf(`🍟: Successfully connected to Redis at ` + address)

	return rdb

}
