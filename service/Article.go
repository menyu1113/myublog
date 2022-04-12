package service

import (
	"myublog/global/consts"
	"myublog/model"
)

func AddArticle(article *model.Article) int {
	err := model.GormDb.Create(&article).Error
	if err != nil {
		return consts.ERRORCODE
	}
	return consts.SUCCSECODE
}
