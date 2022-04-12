package main

import (
	"log"
	"myublog/model"
	"myublog/router"
)

func main() {
	// 初始化数据库
	model.InitDb()
	// 初始化路由
	router.InitRouter()
	log.Printf("myublog开始启动...\n")
}
