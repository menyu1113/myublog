package v1

import (
	"github.com/gin-gonic/gin"
	"myublog/global/consts"
	"myublog/service"
	"myublog/utils/response"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := service.UpLoadFile(file, fileSize)
	response.Response(c, code, consts.GetErrMsg(code), url)

}
