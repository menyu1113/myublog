package service

import (
	"golang.org/x/crypto/bcrypt"
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
)

func CheckLogin(username string, password string) int {
	//if CheckUser(user.Username){//用户不存在
	//	return myerrors.CurdSelectFailCode
	//}
	//return myerrors.CurdStatusOkCode
	var user model.User
	global.GormDb.Model(&model.User{}).Where("username = ?", username).First(&user)
	//用户不存在
	if user.ID == 0 {
		return myerrors.LoginUserNotExistCode
	}
	//密码错误
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return myerrors.FailPasswordCode
	}
	return myerrors.SUCCSECODE

}
