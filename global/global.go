package global

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"todo/config"
)

var (
	DB     *gorm.DB
	LOG    *zap.Logger
	VP     *viper.Viper
	CONFIG config.Config
)
