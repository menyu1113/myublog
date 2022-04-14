package service

import (
	"myublog/global/consts"
	"myublog/middleware/ZapLog"
	"myublog/model"
)
//添加文章
func AddArticle(id uint,article *model.Article) int {
	//根据ID绑定用户
	var user model.User
	err =model.GormDb.First(&user,id).Error
	if err != nil {
		return consts.ERRORCODE
	}
	//创建文章
	err := model.GormDb.Create(article).Error
	if err != nil {
		return consts.ERRORCODE
	}
	//绑定关联
	err = model.GormDb.Model(&user).Association("Articles").Append(article)
	//err = model.GormDb.Model(&article).Association("Users").Append(&user)
	if err != nil {
		ZapLog.ZapLogger.Error("用户文章表关联数据失败:"+err.Error())
		_= model.GormDb.Delete(&article).Error
		return consts.FailUserLinkDetailsCode
	}
	return consts.SUCCSECODE
}
//获取用户单个文章
func GetArtInfo(userid,artid int)(code int,article *model.Article ) {
	err:=model.GormDb.Where("id=? AND user_id=?",artid,userid).First(&article).Error
	if err!=nil{
		ZapLog.ZapLogger.Error("获取用户单个文章失败:"+err.Error())
		return consts.FailGetArticleCode,&model.Article{}
	}
	return consts.SUCCSECODE,article
}
//获取用户所有文章
func GetArt(id uint) (code int,article []model.Article) {
	var user model.User
	err:=model.GormDb.Preload("Articles").First(&user).Error
	if err!=nil{
		ZapLog.ZapLogger.Error("获取用户所有文章失败:"+err.Error())
		return consts.FailGetArticleCode,nil
	}
	return consts.SUCCSECODE,user.Articles
}
// 编辑文章
func EditArt(userid,artid int, article *model.Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img

	err = model.GormDb.Model(&article).Where("id=? AND user_id=?",artid,userid).Updates(&maps).Error
	if err != nil {
		ZapLog.ZapLogger.Error("编辑文章失败:"+err.Error())
		return consts.FailUpdataArticleCode
	}
	return consts.SUCCSECODE
}
//  删除文章
func DeleteArt(artid int) int {
	var art model.Article
	err = model.GormDb.Where("id = ? ", artid).Delete(&art).Error
	if err != nil {
		ZapLog.ZapLogger.Error("删除文章失败:"+err.Error())
		return consts.FailDeleArticleCode
	}
	return consts.SUCCSECODE
}

