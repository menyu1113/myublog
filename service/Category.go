package service

import (
	"gorm.io/gorm"
	"myublog/global/consts"
	"myublog/middleware/ZapLog"
	"myublog/model"
)
// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	if name==""{
		return consts.FailUpdataCateCode //2001
	}
	var cate model.Category
	model.GormDb.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return consts.FailCatenameUsedCode //2001
	}
	return consts.SUCCSECODE
}
func CreateCate(id int,category *model.Category)(code int){
	var user model.User
	err:=model.GormDb.First(&user,id).Error
	if err!=nil{
		ZapLog.ZapLogger.Error("用户不存在:"+err.Error())
		return consts.ERRORCODE
	}
	err=model.GormDb.Create(&category).Error
	if err!=nil{
		ZapLog.ZapLogger.Error("创建分类失败:"+err.Error())
		return consts.ERRORCODE
	}
	//创建用户和分类的关联
	err = model.GormDb.Model(&user).Association("Categorys").Append(category)
	if err != nil {
		ZapLog.ZapLogger.Error("用户分类表关联失败:"+err.Error())
		return  consts.FailUserLinkDetailsCode
	}
	return consts.SUCCSECODE
}

//获取分类
func GetCate(pageSize int, pageNum int) ([]model.Category,int, int64) {
	var cate []model.Category
	var total int64
	err = model.GormDb.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	model.GormDb.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		ZapLog.ZapLogger.Error("获取分类失败:"+err.Error())
		return nil, consts.FailGetCategoryCode,0
	}
	return cate, consts.SUCCSECODE,total
}

// EditCate 编辑分类信息
func EditCate(cateid int, category *model.Category) int {
	var mp = make(map[string]interface{})
	mp["name"] = category.Name
	err = model.GormDb.Model(&category).Where("id = ? ", cateid).Updates(mp).Error
	if err != nil {
		ZapLog.ZapLogger.Error("编辑分类失败:"+err.Error())
		return consts.FailUpdataCateCode
	}
	return consts.SUCCSECODE
}

// DeleteCate 删除分类
func DeleteCate(id,cateid int) int {
	var category model.Category
	err = model.GormDb.Where("id = ? ", cateid).Delete(&category).Error
	if err != nil {
		return consts.FailDeleCateCode
	}
	var user model.User
	model.GormDb.First(&user,id)
	err =model.GormDb.Model(&user).Association("Categorys").Delete(&category)
	if err != nil {
		ZapLog.ZapLogger.Error("删除分类失败:"+err.Error())
		return consts.FailDeleCateCode
	}
	return consts.SUCCSECODE
}

