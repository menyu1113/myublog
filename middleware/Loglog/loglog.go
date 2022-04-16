package Loglog

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myublog/global"
	"time"
)



func ZapLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		Timebegin:=time.Now()
		global.ZapLogger.Info("开始访问",zap.String("访问路径", c.Request.RequestURI),
			zap.String("ip-port",c.Request.Host),
			zap.String("method", c.Request.Method),
			)
		Accessduration:=time.Since(Timebegin)
		c.Next()
		global.ZapLogger.Info("访问结束",zap.String("访问时长为:",Accessduration.String()))
	}
}
