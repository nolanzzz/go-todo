package core

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"todo/global"
	"todo/service"
)

func InitScheduler() {
	c := cron.New()
	global.LOG.Info("cron jobs initialized")
	err := c.AddFunc("@daily", func() { // midnight - 00***
		global.LOG.Info("run CleanRankings")
		service.RankingServiceApp.CleanRankings()
	})
	if err != nil {
		global.LOG.Error("adding cron job failed", zap.Error(err))
		return
	}
	go c.Start()
}
