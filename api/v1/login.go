package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"myublog/global/consts"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"myublog/utils/utiljwt"
	"time"
)

func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("参数绑定失败:%s\n", err)
	}
	code := service.CheckLogin(user.Username, user.Password)
	if code != consts.SUCCSECODE {
		response.Fail(c, code, consts.GetErrMsg(code))
		//response.Response(c,code,consts.GetErrMsg(code))
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		log.Printf("生成token fail%s\n", err)
		return
	}
	//c.SetCookie("token",token ,3600, "/", "localhost", false, true)
	response.Success(c, consts.GetErrMsg(code), token)
}

//生成token令牌
func GenerateToken(user model.User) (string, error) {
	claims := utiljwt.MyClaims{
		Username: user.Username,
		//Id: user.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 1209600,
			Issuer:    "myu",
		},
	}
	j := utiljwt.NewJWT()
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
