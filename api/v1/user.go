package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"myublog/utils/utiljwt"
	"myublog/utils/validators"
	"strconv"
	"strings"
)

//用户注册
func UserRegister(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("参数绑定失败:%s\n", err)
	}
	msg, validCode := validators.Verification(&user)
	if validCode != myerrors.SUCCSECODE {
		response.Fail(c, validCode, msg)
		c.Abort()
		return
	}
	//查询当前用户名是否存在
	if !service.CheckUser(user.Username) { //用户已存在
		response.Fail(c, myerrors.CurdCreatFailExistCode, myerrors.CurdCreatFailExist, "")
		return
	}
	if service.CreateUser(&user) == myerrors.CurdCreatFailCode {
		response.Fail(c, myerrors.CurdCreatFailCode, myerrors.CurdCreatFailMsg, "")
	} else {
		response.Success(c, myerrors.SUCCSE)
	}

}

//查询用户列表
func QueryUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	data, total := service.GetUsers(username, pageSize, pageNum)
	response.Success(c, myerrors.SUCCSE, data, total)
}

//编辑用户
func EditUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("参数绑定失败:%s\n", err)
	}
	code := service.EditUser(id, &user)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.GetErrMsg(code))
		return
	}
	//更改信息成功后，重新生成token并返回

	response.Success(c, myerrors.GetErrMsg(code))
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteUser(id)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.GetErrMsg(code))
	}
	//添加token至黑名单
	myclaims, _ := c.Get("myclaims")
	TokenString := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	err := service.JoinBlackList(TokenString, myclaims.(*utiljwt.MyClaims))
	if err != nil {
		global.ZapLogger.Error("添加token黑名单失败")
		return
	}
	global.ZapLogger.Info("添加token黑名单成功!")
	response.Success(c, myerrors.GetErrMsg(code))
}

//修改密码
func ChangeUserPassword(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	//err := c.ShouldBindBodyWith(&user,binding.JSON)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("user:", user)
		response.Fail(c, myerrors.BindParameterCode, myerrors.GetErrMsg(myerrors.BindParameterCode))
		return
	}
	code := service.ChangePassword(id, &user)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.PasswordChangeFailed)
		return
	}
	response.Response(c, code, myerrors.GetErrMsg(code))
}
