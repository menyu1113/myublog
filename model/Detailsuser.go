package model

import "gorm.io/gorm"

type DetailsUser struct {
	gorm.Model
	UserID int
	User   User
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	Wechat string `gorm:"type:varchar(100)" json:"wechat"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Img    string `gorm:"type:varchar(200)" json:"img"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
}
