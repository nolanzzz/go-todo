package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"todo/config"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	LOG    *zap.Logger
	REDIS  *redis.Client
	CONFIG config.Config
)
