package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"myublog/global/consts"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"myublog/utils/validators"
	"strconv"
)

//用户注册
func UserRegister(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("参数绑定失败:%s\n", err)
	}
	msg, validCode := validators.Verification(&user)
	if validCode != consts.SUCCSECODE {
		response.Fail(c, validCode, msg)
		c.Abort()
		return
	}
	//查询当前用户名是否存在
	if !service.CheckUser(user.Username) { //用户已存在
		response.Fail(c, consts.CurdCreatFailExistCode, consts.CurdCreatFailExist, "")
		return
	}
	if service.CreateUser(&user) == consts.CurdCreatFailCode {
		response.Fail(c, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
	} else {
		response.Success(c, consts.SUCCSE)
	}

}

//查询用户列表
func QueryUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	data, total := service.GetUsers(username, pageSize, pageNum)
	response.Success(c, consts.SUCCSE, data, total)
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
	if code != consts.SUCCSECODE {
		response.Fail(c, code, consts.GetErrMsg(code))
		return
	}
	//更改信息成功后，重新生成token并返回

	response.Success(c, consts.GetErrMsg(code))
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteUser(id)
	response.Response(c, code, consts.GetErrMsg(code))
}

func ChangeUserPassword(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	//err := c.ShouldBindBodyWith(&user,binding.JSON)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("user:", user)
		response.Fail(c, consts.BindParameterCode, consts.GetErrMsg(consts.BindParameterCode))
		return
	}
	code := service.ChangePassword(id, &user)
	if code != consts.SUCCSECODE {
		response.Fail(c, code, consts.PasswordChangeFailed)
		return
	}
	response.Response(c, code, consts.GetErrMsg(code))
}
