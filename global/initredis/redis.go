package initredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"myublog/global"
	"myublog/global/vipers"
	"os"
	"time"
)

func InitRedis() {
	fmt.Printf("%T,%v\n", time.Duration(vipers.RedisFailDialTimeout)*1000000000, time.Duration(vipers.RedisFailDialTimeout)*1000000000)
	rdb := redis.NewClient(&redis.Options{
		Addr:         vipers.RedisAddr,
		DB:           vipers.RedisIndexDb,                                     // use default DB
		Password:     vipers.RedisPassword,                                    // no password set
		DialTimeout:  time.Duration(vipers.RedisFailDialTimeout * 1000000000), //time.Duration()强转为纳秒
		MinIdleConns: vipers.RedisMinIdleConns,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.ZapLogger.Error("连接redis失败:" + err.Error())
		fmt.Println("连接redis失败:" + err.Error() + pong)
		os.Exit(1)
	}
	global.GoRedis = rdb
	global.ZapLogger.Info("连接redis成功!")
}
