package service

import (
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
)
//增加评论
func AddComment(id, articleid int, comment *model.Comment) int {
	var user model.User
	var aricle model.Article
	if global.GormDb.First(&user, id).Error != nil || global.GormDb.First(&aricle, articleid).Error != nil {
		global.ZapLogger.Error("用户或者文章不存在")
		return myerrors.ERRORCODE
	}
	err := global.GormDb.Create(&comment).Error
	if err != nil {
		global.ZapLogger.Error("新增评论失败:" + err.Error())
		return myerrors.ERRORCODE
	}
	//添加关联
	err = global.GormDb.Model(&user).Association("Comments").Append(comment)
	if err != nil {
		global.ZapLogger.Error("用户评论表关联数据失败:" + err.Error())
		_ = global.GormDb.Delete(&comment).Error
		return myerrors.FailUserLinkDetailsCode
	}
	err = global.GormDb.Model(&aricle).Association("Comments").Append(comment)
	if err != nil {
		global.ZapLogger.Error("文章评论表关联数据失败:" + err.Error())
		_ = global.GormDb.Delete(&comment).Error
		return myerrors.FailUserLinkDetailsCode
	}
	return myerrors.SUCCSECODE
}

//删除评论
func DeleteComment(id uint) int {
	var comment model.Comment
	//comment的id应该要提前判断一下  这里没有判断，结果就是没有删除，但是返回还是成功的
	err = global.GormDb.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return myerrors.FailDeleCommentCode
	}
	return myerrors.SUCCSECODE
}

//获取文章所有
func GetComment(id uint) (code int,comment []model.Comment) {
	var article model.Article
	err:= global.GormDb.Preload("Comments").First(&article,"id = ?",id).Error
	if err!=nil{
		global.ZapLogger.Error("获取文章所有失败:"+err.Error())
		return myerrors.FailGetCommentCode,nil
	}
	return myerrors.SUCCSECODE,article.Comments
}