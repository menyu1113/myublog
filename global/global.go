package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GormDb     *gorm.DB
	GoRedis  *redis.Client
	ZapLogger *zap.Logger
	Sugar     *zap.SugaredLogger
)

