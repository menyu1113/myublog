package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
)

var err error

//创建用户
func CreateUser(user *model.User) int {
	user.Password, err = encrypTion(user.Password)
	if err != nil {
		global.ZapLogger.Error("密码加密失败:"+err.Error())
		return myerrors.CurdCreatFailCode
	}
	err = global.GormDb.Create(&user).Error
	if err != nil {
		global.ZapLogger.Error("创建用户失败:"+err.Error())
		return myerrors.CurdCreatFailCode
	}
	return myerrors.SUCCSECODE
}

//密码散列
func encrypTion(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

// CheckUser 检查用户是否存在 false表示存在
func CheckUser(name string) bool {
	var user model.User
	//用户存在，Error如果为nil 那么errors.Is就是false, 取反就是不存在
	if !errors.Is(global.GormDb.Where("username = ?", name).First(&user).Error, gorm.ErrRecordNotFound) {
		global.ZapLogger.Error(gorm.ErrRecordNotFound.Error())
		return false //（不存在返回false）
	}
	return true
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]model.User, int64) {

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	var users []model.User
	var total int64

	if username != "" {
		global.GormDb.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		global.GormDb.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	global.GormDb.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	global.GormDb.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

// 编辑用户
func EditUser(id int, user *model.User) int {
	//查询用户名是否与已有重复或者id是否存在
	code := CheckUerAndId(id, user.Username)
	if code != myerrors.SUCCSECODE {
		return code
	}

	var mp = make(map[string]interface{})
	mp["username"] = user.Username
	mp["role"] = user.Role
	err = global.GormDb.Model(&user).Where("id = ?", id).Updates(mp).Error
	if err != nil {
		global.ZapLogger.Error("更新用户失败:"+err.Error())
		return myerrors.CurdUpdateFailCode
	}
	return myerrors.SUCCSECODE
}
func CheckUerAndId(id int, name string) (code int) {
	var user model.User
	//判断ID是否存在 id不存在，Error有值 等于gorm.ErrRecordNotFound，errors.Is为true
	if errors.Is(global.GormDb.First(&user, id).Error, gorm.ErrRecordNotFound) {
		return myerrors.CurdFailNotExistCode
	}

	//err=model.GormDb.First(&user,id).Error
	//if err !=nil{
	//	return myerrors.CurdFailNotExistCode
	//}
	user= model.User{}
	//用户是否存在
	global.GormDb.Select("id, username").Where("username = ?", name).First(&user)
	//fmt.Printf("name:%v, id:%v, user.Username:%v,user.ID:%v\n",name,id,user.Username,user.ID)
	if user.ID == uint(id) {
		return myerrors.SUCCSECODE
	}
	if user.ID > 0 { //用户名已存在，不能用这个用户名了
		return myerrors.CurdUpdateFailExistCode
	}
	return myerrors.SUCCSECODE
}

//删除用户
func DeleteUser(id int) (code int) {
	var user model.User
	err = global.GormDb.Model(&user).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		global.ZapLogger.Error("删除用户失败:"+err.Error())
		return myerrors.CurdDeleteFailCode
	}
	return myerrors.SUCCSECODE
}
func ChangePassword(id int, user *model.User) (code int) {
	//查传过来的ID是不是当前用户ID，不是不能改
	var temp model.User
	global.GormDb.Select("username").Where("id = ?", id).First(&temp)
	if temp.Username != user.Username {
		return myerrors.ERRORCODE
	}
	user.Password, _ = encrypTion(user.Password) //加密
	err = global.GormDb.Select("password").Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return myerrors.ERRORCODE
	}
	return myerrors.SUCCSECODE
}
