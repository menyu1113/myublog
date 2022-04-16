package service

import (
	"gorm.io/gorm"
	"myublog/global"
	"myublog/global/myerrors"
	"myublog/model"
)
// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	if name==""{
		return myerrors.FailUpdataCateCode //2001
	}
	var cate model.Category
	global.GormDb.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return myerrors.FailCatenameUsedCode //2001
	}
	return myerrors.SUCCSECODE
}
func CreateCate(id int,category *model.Category)(code int){
	var user model.User
	err:= global.GormDb.First(&user,id).Error
	if err!=nil{
		global.ZapLogger.Error("用户不存在:"+err.Error())
		return myerrors.ERRORCODE
	}
	err= global.GormDb.Create(&category).Error
	if err!=nil{
		global.ZapLogger.Error("创建分类失败:"+err.Error())
		return myerrors.ERRORCODE
	}
	//创建用户和分类的关联
	err = global.GormDb.Model(&user).Association("Categorys").Append(category)
	if err != nil {
		global.ZapLogger.Error("用户分类表关联失败:"+err.Error())
		return  myerrors.FailUserLinkDetailsCode
	}
	return myerrors.SUCCSECODE
}

//获取分类
func GetCate(pageSize int, pageNum int) ([]model.Category,int, int64) {
	var cate []model.Category
	var total int64
	err = global.GormDb.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	global.GormDb.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.ZapLogger.Error("获取分类失败:"+err.Error())
		return nil, myerrors.FailGetCategoryCode,0
	}
	return cate, myerrors.SUCCSECODE,total
}

// EditCate 编辑分类信息
func EditCate(cateid int, category *model.Category) int {
	var mp = make(map[string]interface{})
	mp["name"] = category.Name
	err = global.GormDb.Model(&category).Where("id = ? ", cateid).Updates(mp).Error
	if err != nil {
		global.ZapLogger.Error("编辑分类失败:"+err.Error())
		return myerrors.FailUpdataCateCode
	}
	return myerrors.SUCCSECODE
}

// DeleteCate 删除分类
func DeleteCate(id,cateid int) int {
	var category model.Category
	err = global.GormDb.Where("id = ? ", cateid).Delete(&category).Error
	if err != nil {
		return myerrors.FailDeleCateCode
	}
	var user model.User
	global.GormDb.First(&user,id)
	err = global.GormDb.Model(&user).Association("Categorys").Delete(&category)
	if err != nil {
		global.ZapLogger.Error("删除分类失败:"+err.Error())
		return myerrors.FailDeleCateCode
	}
	return myerrors.SUCCSECODE
}

