package middleware

import (
	"github.com/gin-gonic/gin"
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"myublog/utils/utiljwt"
	"strings"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			response.Fail(c, myerrors.ERRORCODE, utiljwt.TokenErrorsExist.Error())
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			response.Fail(c, myerrors.ERRORCODE, utiljwt.TokenNotValidYet.Error())
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			response.Fail(c, myerrors.ERRORCODE, utiljwt.TokenNotValidYet.Error())
			c.Abort()
			return
		}
		//判断token是否在黑名单中
		if service.IsInBlacklist(checkToken[1]) {
			response.Fail(c, myerrors.ERRORCODE, utiljwt.TokenBlackList.Error())
			c.Abort()
			return
		}
		j := utiljwt.NewJWT()
		// 解析token
		myclaims, err := j.ParserToken(checkToken[1])
		if err != nil {
			response.Fail(c, myerrors.ERRORCODE, err.Error())
			c.Abort()
			return
		}
		//检查token是不是当前用户的
		//id, _ := strconv.Atoi(c.Param("id"))
		//if !check(id, myclaims.Username) {
		//	response.Fail(c, myerrors.ERRORCODE, utiljwt.TokenNoCurrentUser.Error())
		//	c.Abort()
		//	return
		//}
		c.Set("myclaims", myclaims)
		c.Next()
	}
}

//检查token是不是当前用户的 不是返回false
func check(id int, name string) bool {
	var user model.User
	global.GormDb.Select("id").Where("username=?", name).First(&user)
	//fmt.Printf("user:%v,user.ID:%v,uint(id):%v\n", user, user.ID, uint(id))
	if user.ID != uint(id) {
		return false
	}
	return true
}
