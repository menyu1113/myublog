package v1

import (
	"github.com/gin-gonic/gin"
	"myublog/global/consts"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"strconv"
)

//获取用户详情
func GetDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.GetDetails(id)
	response.Response(c, code, consts.GetErrMsg(code), data)
}

//更新用户详情
func UpdateDetails(c *gin.Context) {
	var data model.DetailsUser
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.UpdateDetails(id, &data)
	response.Response(c, code, consts.GetErrMsg(code))
}
//添加用户详情
func AddUserDetail(c *gin.Context){
	var detail model.DetailsUser
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&detail)
	code:=service.AddUserDetail(&detail,id)
	if code !=consts.SUCCSECODE{
		response.Fail(c,code,consts.GetErrMsg(code))
		return
	}
	response.Success(c,consts.GetErrMsg(code))
}
