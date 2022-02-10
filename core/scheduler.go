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
	var err error
	// Clean rankings in Redis at midgnight everyday
	if err = c.AddFunc("@daily", func() {
		global.LOG.Info("run CleanRankings")
		service.RankingServiceApp.CleanRankings()
	}); err != nil {
		global.LOG.Error("adding cron job failed", zap.Error(err), zap.String("job", "CleanRankings"))
		return
	}

	// Generate 5 random todos every 5 minutes
	if err = c.AddFunc("0 */5 * * *", func() {
		global.LOG.Info("run GenerateTodos")
		service.TodoServiceApp.GenerateTodos(5)
	}); err != nil {
		global.LOG.Error("adding cron job failed", zap.Error(err), zap.String("job", "GenerateTodos"))
		return
	}

	// Complete todos from a random user every 10 minutes
	if err = c.AddFunc("0 */10 * * *", func() {
		global.LOG.Info("run GenerateTodos")
		service.TodoServiceApp.CompleteTodos(5)
	}); err != nil {
		global.LOG.Error("adding cron job failed", zap.Error(err), zap.String("job", "CompleteTodos"))
		return
	}

	go c.Start()

	//t1 := time.NewTimer(time.Second * 5)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 5)
	//	}
	//}
}
