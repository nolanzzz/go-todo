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
		panic("Failed to connect to Redis: " + err.Error())
	} else {
		global.LOG.Info("redis connect ping response", zap.String("pong", pong))
	}
	if err = cleanExistingData(client); err != nil {
		global.LOG.Error("reset redis database failed", zap.Error(err))
		panic("Failed to reset Redis: " + err.Error())
	}
	return client
}

func cleanExistingData(client *redis.Client) (err error) {
	err = client.Del(global.CONFIG.Redis.KeyRankTodos, global.CONFIG.Redis.KeyRankMinutes).Err()
	return err
}
