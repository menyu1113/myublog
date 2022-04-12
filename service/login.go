package service

import (
	"golang.org/x/crypto/bcrypt"
	"myublog/global/consts"
	"myublog/model"
)

func CheckLogin(username string, password string) int {
	//if CheckUser(user.Username){//用户不存在
	//	return consts.CurdSelectFailCode
	//}
	//return consts.CurdStatusOkCode
	var user model.User
	model.GormDb.Model(&model.User{}).Where("username = ?", username).First(&user)
	//用户不存在
	if user.ID == 0 {
		return consts.LoginUserNotExistCode
	}
	//密码错误
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return consts.ERROR_PASSWORD_Code
	}
	return consts.SUCCSECODE

}
