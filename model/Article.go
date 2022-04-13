package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	CategoryID uint
	UserID uint
	Comments []Comment
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}
