package v1

import (
	"github.com/gin-gonic/gin"
	"myublog/global/consts"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"strconv"
)

// AddComment 新增评论
func AddComment(c *gin.Context) {
	var comment model.Comment
	id, err1 := strconv.Atoi(c.Param("id"))
	articleid,err2:=strconv.Atoi(c.Query("articleid"))
	if err2 !=nil || err1 !=nil{
		response.Fail(c, consts.ERRORCODE, consts.ERROR)
		return
	}
	_ = c.ShouldBindJSON(&comment)
	code := service.AddComment(id,articleid,&comment)
	response.Response(c, code, consts.GetErrMsg(code))
}
// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DeleteComment(uint(id))
	if code!=consts.SUCCSECODE{
		response.Fail(c,code,consts.GetErrMsg(code))
		return
	}
	response.Success(c,consts.GetErrMsg(code))
}
//获取文章所有评论
func GetComment(c *gin.Context)  {
	artid,_:=strconv.Atoi(c.Query("artid"))
	code,comment:=service.GetComment(uint(artid))
	if code!=consts.SUCCSECODE{
		response.Fail(c,code,consts.GetErrMsg(code))
		return
	}
	response.Success(c,consts.SUCCSE,comment)

}