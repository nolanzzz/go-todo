package service

import (
	"github.com/go-redis/redis"
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
