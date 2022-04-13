package model

type Blog struct {
	ID        uint `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint
	SiteName  string `gorm:"type:varchar(100);not null" json:"sitename" label:"站点名称"`
	SiteDesc  string `gorm:"type:varchar(200)" json:"sitedesc" label:"站点描述"`
	SiteTitle string `gorm:"type:varchar(200)" json:"sitetitle" label:"站点标题"`
	SiteTheme string `gorm:"type:longtext" json:"sitetheme" label:"站点样式"`
}
