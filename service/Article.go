package service

import (
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
)
//添加文章
func AddArticle(id uint,article *model.Article) int {
	//根据ID绑定用户
	var user model.User
	err = global.GormDb.First(&user,id).Error
	if err != nil {
		return myerrors.ERRORCODE
	}
	//创建文章
	err := global.GormDb.Create(article).Error
	if err != nil {
		return myerrors.ERRORCODE
	}
	//绑定关联
	err = global.GormDb.Model(&user).Association("Articles").Append(article)
	//err = model.GormDb.Model(&article).Association("Users").Append(&user)
	if err != nil {
		global.ZapLogger.Error("用户文章表关联数据失败:"+err.Error())
		_= global.GormDb.Delete(&article).Error
		return myerrors.FailUserLinkDetailsCode
	}
	return myerrors.SUCCSECODE
}
//获取用户单个文章
func GetArtInfo(userid,artid int)(code int,article *model.Article ) {
	err:= global.GormDb.Where("id=? AND user_id=?",artid,userid).First(&article).Error
	if err!=nil{
		global.ZapLogger.Error("获取用户单个文章失败:"+err.Error())
		return myerrors.FailGetArticleCode,&model.Article{}
	}
	return myerrors.SUCCSECODE,article
}
//获取用户所有文章
func GetArt(id uint) (code int,article []model.Article) {
	var user model.User
	err:= global.GormDb.Preload("Articles").First(&user).Error
	if err!=nil{
		global.ZapLogger.Error("获取用户所有文章失败:"+err.Error())
		return myerrors.FailGetArticleCode,nil
	}
	return myerrors.SUCCSECODE,user.Articles
}
// 编辑文章
func EditArt(userid,artid int, article *model.Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img

	err = global.GormDb.Model(&article).Where("id=? AND user_id=?",artid,userid).Updates(&maps).Error
	if err != nil {
		global.ZapLogger.Error("编辑文章失败:"+err.Error())
		return myerrors.FailUpdataArticleCode
	}
	return myerrors.SUCCSECODE
}
//  删除文章
func DeleteArt(artid int) int {
	var art model.Article
	err = global.GormDb.Where("id = ? ", artid).Delete(&art).Error
	if err != nil {
		global.ZapLogger.Error("删除文章失败:"+err.Error())
		return myerrors.FailDeleArticleCode
	}
	return myerrors.SUCCSECODE
}

