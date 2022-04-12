package v1

import (
	"github.com/gin-gonic/gin"
	"myublog/global/consts"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	code := service.AddArticle(&article)
	response.Response(c, code, consts.GetErrMsg(code))
}
