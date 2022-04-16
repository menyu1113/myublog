package main

import (
	"log"
	"myublog/global/gorm"
	"myublog/global/initredis"
	"myublog/router"
	"myublog/utils/zaplog"
)

func main() {
	//初始化日志
	zaplog.InitZapLog()
	//初始化mysql数据库
	gorm.InitDb()
	//初始化redis
	initredis.InitRedis()
	// 初始化路由
	router.InitRouter()
	log.Printf("myublog开始启动...\n")
}
