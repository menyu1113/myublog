package ZapLog

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"myublog/global/vipers"
	"os"
	"strings"
	"time"
)

var (
	err       error
	ZapLogger *zap.Logger
	Sugar     *zap.SugaredLogger
)

func ZapLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if vipers.AppDebug == true {
		//	if ZapLogger, err = zap.NewDevelopment(); err == nil {
		//		return
		//	} else {
		//		log.Fatal("创建zap日志包失败，详情：" + err.Error())
		//	}
		//}
		lumberjackhook := lumberjack.Logger{
			Filename:   vipers.Filename,
			MaxSize:    vipers.MaxSize,
			MaxAge:     vipers.MaxAge,
			MaxBackups: vipers.MaxBackups,
			Compress:   vipers.Compress,
		}
		write := zapcore.AddSync(&lumberjackhook)

		encoderConfig := zap.NewProductionEncoderConfig()
		//日志输入格式
		//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		var encoder zapcore.Encoder
		switch strings.ToLower(vipers.TextLogFormat) {
		case "json":
			encoder = zapcore.NewJSONEncoder(encoderConfig)
		case "console":
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		default:
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		}

		core := zapcore.NewCore(
			encoder,
			write,
			zapcore.InfoLevel,
		)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		filed := zap.Fields(
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.RequestURI),
			zap.String("Ip-Port", c.Request.Host),
			zap.String("hostName", hostName),
		)
		ZapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel), filed)
		Sugar = ZapLogger.Sugar()
		ZapLogger.Info("")
	}

}
