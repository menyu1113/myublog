package service

import (
	"myublog/global/consts"
	"myublog/middleware/ZapLog"
	"myublog/model"
)
//增加评论
func AddComment(id, articleid int, comment *model.Comment) int {
	var user model.User
	var aricle model.Article
	if model.GormDb.First(&user, id).Error != nil || model.GormDb.First(&aricle, articleid).Error != nil {
		ZapLog.ZapLogger.Error("用户或者文章不存在")
		return consts.ERRORCODE
	}
	err := model.GormDb.Create(&comment).Error
	if err != nil {
		ZapLog.ZapLogger.Error("新增评论失败:" + err.Error())
		return consts.ERRORCODE
	}
	//添加关联
	err = model.GormDb.Model(&user).Association("Comments").Append(comment)
	if err != nil {
		ZapLog.ZapLogger.Error("用户评论表关联数据失败:" + err.Error())
		_ = model.GormDb.Delete(&comment).Error
		return consts.FailUserLinkDetailsCode
	}
	err = model.GormDb.Model(&aricle).Association("Comments").Append(comment)
	if err != nil {
		ZapLog.ZapLogger.Error("文章评论表关联数据失败:" + err.Error())
		_ = model.GormDb.Delete(&comment).Error
		return consts.FailUserLinkDetailsCode
	}
	return consts.SUCCSECODE
}

//删除评论
func DeleteComment(id uint) int {
	var comment model.Comment
	//comment的id应该要提前判断一下  这里没有判断，结果就是没有删除，但是返回还是成功的
	err = model.GormDb.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return consts.FailDeleCommentCode
	}
	return consts.SUCCSECODE
}

//获取文章所有
func GetComment(id uint) (code int,comment []model.Comment) {
	var article model.Article
	err:=model.GormDb.Preload("Comments").First(&article,"id = ?",id).Error
	if err!=nil{
		ZapLog.ZapLogger.Error("获取文章所有失败:"+err.Error())
		return consts.FailGetCommentCode,nil
	}
	return consts.SUCCSECODE,article.Comments
}