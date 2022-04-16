package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"myublog/global/myerrors"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"myublog/utils/utiljwt"
)

func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("参数绑定失败:%s\n", err)
	}
	code := service.CheckLogin(user.Username, user.Password)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.GetErrMsg(code))
		//response.Response(c,code,myerrors.GetErrMsg(code))
		return
	}
	token, err := utiljwt.GenerateToken(user)
	if err != nil {
		log.Printf("生成token fail%s\n", err)
		return
	}
	//c.SetCookie("token",token ,3600, "/", "localhost", false, true)
	response.Success(c, myerrors.GetErrMsg(code), token)
}
