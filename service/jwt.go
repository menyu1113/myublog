package service

import (
	"context"
	"myublog/global"
	"myublog/utils/md5x"
	"myublog/utils/utiljwt"
	"time"
)

// 获取黑名单缓存 key
func getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + md5x.MD5([]byte(tokenStr))
}

// JoinBlackList token 加入黑名单
func JoinBlackList(TokenString string, token *utiljwt.MyClaims) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.ExpiresAt-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	err = global.GoRedis.SetNX(context.Background(), getBlackListKey(TokenString), nowUnix, timer).Err()
	return
}

// IsInBlacklist token 是否在黑名单中
func IsInBlacklist(TokenString string) bool {
	joinUnixStr, err := global.GoRedis.Get(context.Background(), getBlackListKey(TokenString)).Result()
	//joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	//if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
	//	return false
	//}
	return true
}
