package response

import (
	"github.com/gin-gonic/gin"
	"myublog/global/consts"
	"net/http"
)

func ResponseJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}
func Success(c *gin.Context, msg string, data ...interface{}) {
	ResponseJson(c, http.StatusOK, consts.SUCCSECODE, msg, data)
}
func Fail(c *gin.Context, dataCode int, msg string, data ...interface{}) {
	ResponseJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

//封装返回参数
func Response(c *gin.Context, dataCode int, msg string, data ...interface{}) {
	if dataCode != consts.SUCCSECODE {
		Success(c, msg, data)
		return
	} else {
		Fail(c, dataCode, msg, data)
	}
}
