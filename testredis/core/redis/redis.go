package redis

import (
	"github.com/go-redis/redis"
	"github.com/op/go-logging"
	"myproj/try/testredis/config"
)

var client *redis.Client
var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("redis")
}

func Init() {
	rediscli := GetClient()
	if rediscli != nil {
		logger.Info("good")
	}
}

func GetClient() *redis.Client {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr: config.Globalconfig.QabircConfig.Redis.Address,
		})
	}
	return client
}
