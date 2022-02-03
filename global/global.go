package global

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	DB  *gorm.DB
	LOG *zap.Logger
)
