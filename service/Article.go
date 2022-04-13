package service

import (
	"myublog/global/consts"
	"myublog/model"
)
//添加文章
func AddArticle(article *model.Article) int {
	err := model.GormDb.Create(&article).Error
	if err != nil {
		return consts.ERRORCODE
	}
	return consts.SUCCSECODE
}
