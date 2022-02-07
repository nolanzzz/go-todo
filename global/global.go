package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"todo/config"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	LOG    *zap.Logger
	REDIS  *redis.Client
	CONFIG config.Config
)
