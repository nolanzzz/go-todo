package service

import (
	"github.com/go-redis/redis"
	"todo/global"
)

type RankingService struct{}

var RankingServiceApp RankingService

func (r *RankingService) RankingByTodos(limit int) ([]redis.Z, error) {
	//var rankings map[string]int
	rankings, err := global.REDIS.ZRangeWithScores(global.CONFIG.Redis.KeyRankTodos, 0, int64(limit)).Result()
	if err != nil {
		return nil, err
	}
	return rankings, nil
}
