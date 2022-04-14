package service

import (
	"myublog/global/consts"
	"myublog/middleware/ZapLog"
	"myublog/model"
)

//获取用户详情
func GetDetails(id int) (model.DetailsUser, int) {
	var Detail model.DetailsUser
	err = model.GormDb.Preload("User").Where("ID = ?", id).First(&Detail).Error
	if err != nil {
		ZapLog.ZapLogger.Error("获取用户详情失败:"+err.Error())
		return Detail, consts.FailGetUserDetailsCode
	}
	return Detail, consts.SUCCSECODE
}

//更新用户详情
func UpdateDetails(id int, data *model.DetailsUser) int {
	var Detail model.DetailsUser
	err = model.GormDb.Model(&Detail).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		ZapLog.ZapLogger.Error("更新用户失败:"+err.Error())
		return consts.FailUpdataUserDetailsCode
	}
	return consts.SUCCSECODE
}

func AddUserDetail(Detail *model.DetailsUser,id int)int {
	//user.Password, err = encrypTion(user.Password)
	//if err != nil {
	//	return consts.CurdCreatFailCode
	//}
	//err = model.GormDb.Create(&user).Error
	var user model.User
	err = model.GormDb.Where("id = ?", id).First(&user).Error
	if err != nil {
		ZapLog.ZapLogger.Error("查询用户失败:"+err.Error())
		return consts.CurdCreatFailCode
	}
	err = model.GormDb.Create(&Detail).Error
	if err != nil {
		ZapLog.ZapLogger.Error("创建文章失败:"+err.Error())
		return consts.CurdCreatFailCode
	}
	//fmt.Printf("Detail:%+v,user:%v,\n",*Detail,&user)
	//err := model.GormDb.Model(&user).Association("DetailsUser").Append(&Detail)//错误示例
	err := model.GormDb.Model(&Detail).Association("User").Append(&user)
	if err != nil {
		ZapLog.ZapLogger.Error("添加用户详情失败:"+err.Error())
		_=model.GormDb.Delete(&Detail)
		return consts.FailUserLinkDetailsCode
	}
	//blog:=model.Blog{SiteName: "站点名1"}
	//model.GormDb.Create(&blog)
	//model.GormDb.Model(&user).Association("Blog").Append(&blog)
	return consts.SUCCSECODE
}
