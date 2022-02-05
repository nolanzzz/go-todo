package core

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"todo/global"
)

func Redis() (client *redis.Client) {
	configs := global.CONFIG.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     configs.Addr,
		Password: configs.Password,
		DB:       configs.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed", zap.Error(err))
		panic("Failed to connect to Redis")
	} else {
		global.LOG.Info("redis connect ping response", zap.String("pong", pong))
	}
	return client
}
