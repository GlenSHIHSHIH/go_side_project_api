package model

import "time"

type Carousel struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"comment:圖片名稱;type:varchar(100)" json:"name"`                                      //圖片名稱
	Image        string    `gorm:"comment:主圖片url;type:text" json:"image"`                                           //主圖片
	Url          string    `gorm:"comment:蝦皮連結;type:text" json:"url"`                                               //蝦皮連結
	Weight       int       `gorm:"comment:權重(優先順序 重=高);type:int" json:"weight"`                                     //蝦皮連結
	CreateTime   time.Time `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime   time.Time `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
	CreateUserId int       `gorm:"comment:新增人員Id;type:int; json:"createUserId"`                                     //新增人員
	UpdateUserId int       `gorm:"comment:修改人員Id;type:int; json:"updateUserId"`                                     //修改人員
}