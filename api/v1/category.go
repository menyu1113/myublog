package v1

import (
	"github.com/gin-gonic/gin"
	"myublog/global/myerrors"
	"myublog/model"
	"myublog/service"
	"myublog/utils/response"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	code := service.CheckCategory(category.Name)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.GetErrMsg(code))
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	code=service.CreateCate(id,&category)
	if code!= myerrors.SUCCSECODE{
		response.Fail(c, code, myerrors.GetErrMsg(code))
		return
	}
	response.Success(c, myerrors.GetErrMsg(code))
}

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := service.GetCate(pageSize, pageNum)
	response.Response(c, code, myerrors.GetErrMsg(code), data, total)
}

// EditCate 编辑分类名
func EditCate(c *gin.Context) {
	var category model.Category
	cateid, _ := strconv.Atoi(c.Param("cateid"))
	_ = c.ShouldBindJSON(&category)
	code := service.CheckCategory(category.Name)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.GetErrMsg(code))
		return
	}
	code =service.EditCate(cateid,&category)
	response.Response(c, code, myerrors.GetErrMsg(code))
}

// DeleteCate 删除用户分类
func DeleteCate(c *gin.Context) {
	id, err2 := strconv.Atoi(c.Param("id"))
	cateid, err1 := strconv.Atoi(c.Query("cateid"))
	if err2 !=nil || err1 !=nil{
		response.Fail(c, myerrors.FailDeleCateCode, myerrors.FailDeleCate)
		return
	}

	code := service.DeleteCate(id,cateid)
	if code != myerrors.SUCCSECODE {
		response.Fail(c, code, myerrors.GetErrMsg(code))
		return
	}
	response.Success(c, myerrors.GetErrMsg(code))
}

