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

// RankingByTodos
// @Tags 	Ranking
// @Summary Get ranking by number of todos
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"ranking":{}},"msg":"succeed"}"
// @Param   limit path int true "number of records"
// @Router 	/api/v1/ranking/todos/:limit [get]
func (r *RankingApi) RankingByTodos(c *gin.Context) {
	limit := c.GetInt("limit")
	if records, err := service.RankingServiceApp.Ranking(limit, "todos", "desc"); err != nil {
		global.LOG.Error("get ranking by todos from redis failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	} else {
		response.OkWithData(c, gin.H{"ranking": records})
	}
}

// RankingByMinutes
// @Tags 	Ranking
// @Summary Get ranking by total sum of minutes
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"ranking":{}},"msg":"succeed"}"
// @Param   limit path int true "number of records"
// @Router 	/api/v1/ranking/minutes/:limit [get]
func (r *RankingApi) RankingByMinutes(c *gin.Context) {
	limit := c.GetInt("limit")
	if records, err := service.RankingServiceApp.Ranking(limit, "minutes", "desc"); err != nil {
		global.LOG.Error("get ranking by minutes from redis failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	} else {
		response.OkWithData(c, gin.H{"ranking": records})
	}
}
