package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"todo/global"
	"todo/model/response"
	"todo/service"
)

type RankingApi struct{}

var Ranking RankingApi

func (r *RankingApi) RankingByTodos(c *gin.Context) {
	limit := c.GetInt("limit")
	if limit < 10 {
		limit = 10 // get top 10 records at least
	}
	if records, err := service.RankingServiceApp.RankingByTodos(limit); err != nil {
		global.LOG.Error("get ranking by todos from redis failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	} else {
		response.OkWithData(c, gin.H{"ranking": records})
	}
}
