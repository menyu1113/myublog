package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Blog Blog
	Comments []Comment //用户:评论 1：N
	Categorys []Category //用户:分类 1：N
	Articles  []Article //用户:文章 1：N
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}
