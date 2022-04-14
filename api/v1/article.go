package v1

import (
	"github.com/gin-gonic/gin"
	"myublog/global/consts"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	code := service.AddArticle(uint(id), &article)
	response.Response(c, code, consts.GetErrMsg(code))
}

//获取用户单个文章信息
func GetArtInfo(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("id"))
	artid, _ := strconv.Atoi(c.Query("artid"))
	code, article := service.GetArtInfo(userid, artid)
	response.Response(c, code, consts.GetErrMsg(code), article)
}

// 查询用户所有文章信息
func GetArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//var article model.Article
	code, article := service.GetArt(uint(id))
	response.Response(c, code, consts.GetErrMsg(code), article)
}

//编辑文章
func EditArt(c *gin.Context) {
	var article model.Article
	userid, _ := strconv.Atoi(c.Param("userid"))
	artid, _ := strconv.Atoi(c.Param("artid"))
	_ = c.ShouldBindJSON(&article)

	code := service.EditArt(userid, artid, &article)

	response.Response(c, code, consts.GetErrMsg(code))
}

// DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	artid, _ := strconv.Atoi(c.Param("artid"))
	code := service.DeleteArt(artid)
	if code != consts.SUCCSECODE {
		response.Fail(c, code, consts.GetErrMsg(code))
		return
	}
	response.Success(c, consts.GetErrMsg(code))
}
