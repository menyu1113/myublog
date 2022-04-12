package service

import (
	"fmt"
	"myublog/global/consts"
	"myublog/model"
)

//获取用户详情
func GetDetails(id int) (model.DetailsUser, int) {
	var Detail model.DetailsUser
	err = model.GormDb.Preload("User").Where("ID = ?", id).First(&Detail).Error
	if err != nil {
		return Detail, consts.FailGetUserDetailsCode
	}
	return Detail, consts.SUCCSECODE
}

//更新用户详情
func UpdateDetails(id int, data *model.DetailsUser) int {
	var Detail model.DetailsUser
	err = model.GormDb.Model(&Detail).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
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
	err = model.GormDb.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		return consts.CurdCreatFailCode
	}
	err = model.GormDb.Debug().Create(&Detail).Error
	if err != nil {
		return consts.CurdCreatFailCode
	}
	fmt.Printf("Detail:%+v,user:%v,\n",*Detail,&user)
	err := model.GormDb.Model(&Detail).Association("User").Append(&user)
	if err != nil {
		return consts.FailUserLinkDetailsCode
	}
	return consts.SUCCSECODE
}
