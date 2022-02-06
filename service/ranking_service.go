package service

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"todo/global"
)

type RankingService struct{}

var RankingServiceApp RankingService

func (r *RankingService) Ranking(limit int, category string) ([]redis.Z, error) {
	if limit < 10 {
		limit = 10 // get top 10 records at least
	}
	var key string
	if category == "todos" {
		key = global.CONFIG.Redis.KeyRankTodos
	} else {
		key = global.CONFIG.Redis.KeyRankMinutes
	}
	rankings, err := global.REDIS.ZRangeWithScores(key, 0, int64(limit)).Result()
	if err != nil {
		return nil, err
	}
	return rankings, nil
}

func (r *RankingService) CleanRankings() {
	if err := global.REDIS.Del(global.CONFIG.Redis.KeyRankMinutes).Err(); err != nil {
		global.LOG.Error("cron - clean redis minutes ranking failed", zap.Error(err))
	}
	if err := global.REDIS.Del(global.CONFIG.Redis.KeyRankTodos).Err(); err != nil {
		global.LOG.Error("cron - clean redis todos ranking failed", zap.Error(err))
	}
}
