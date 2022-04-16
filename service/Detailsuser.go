package service

import (
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
)

//获取用户详情
func GetDetails(id int) (model.DetailsUser, int) {
	var Detail model.DetailsUser
	err = global.GormDb.Preload("User").Where("ID = ?", id).First(&Detail).Error
	if err != nil {
		global.ZapLogger.Error("获取用户详情失败:"+err.Error())
		return Detail, myerrors.FailGetUserDetailsCode
	}
	return Detail, myerrors.SUCCSECODE
}

//更新用户详情
func UpdateDetails(id int, data *model.DetailsUser) int {
	var Detail model.DetailsUser
	err = global.GormDb.Model(&Detail).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		global.ZapLogger.Error("更新用户失败:"+err.Error())
		return myerrors.FailUpdataUserDetailsCode
	}
	return myerrors.SUCCSECODE
}

func AddUserDetail(Detail *model.DetailsUser,id int)int {
	//user.Password, err = encrypTion(user.Password)
	//if err != nil {
	//	return myerrors.CurdCreatFailCode
	//}
	//err = model.GormDb.Create(&user).Error
	var user model.User
	err = global.GormDb.Where("id = ?", id).First(&user).Error
	if err != nil {
		global.ZapLogger.Error("查询用户失败:"+err.Error())
		return myerrors.CurdCreatFailCode
	}
	err = global.GormDb.Create(&Detail).Error
	if err != nil {
		global.ZapLogger.Error("创建文章失败:"+err.Error())
		return myerrors.CurdCreatFailCode
	}
	//fmt.Printf("Detail:%+v,user:%v,\n",*Detail,&user)
	//err := model.GormDb.Model(&user).Association("DetailsUser").Append(&Detail)//错误示例
	err := global.GormDb.Model(&Detail).Association("User").Append(&user)
	if err != nil {
		global.ZapLogger.Error("添加用户详情失败:"+err.Error())
		_= global.GormDb.Delete(&Detail)
		return myerrors.FailUserLinkDetailsCode
	}
	//blog:=model.Blog{SiteName: "站点名1"}
	//model.GormDb.Create(&blog)
	//model.GormDb.Model(&user).Association("Blog").Append(&blog)
	return myerrors.SUCCSECODE
}
