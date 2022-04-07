package model

import (
	"time"

	"gorm.io/gorm"
)

type Carousel struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"comment:圖片名稱;type:varchar(100)" json:"name"`                                      //圖片名稱
	Weight       int            `gorm:"comment:權重(優先順序 重=高);type:int;default:0" json:"weight"`                           //權重
	Status       *bool          `gorm:"comment:開關 (true=開啟);type:bool;default:true" json:"status"`                       //狀態(開關)
	StartTime    time.Time      `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"startTime"`  //開始時間
	EndTime      time.Time      `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"endTime"`    //結束時間
	CreateTime   time.Time      `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime   time.Time      `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
	CreateUserId int            `gorm:"comment:新增人員Id;type:int; json:"createUserId"`                                     //新增人員
	UpdateUserId int            `gorm:"comment:修改人員Id;type:int; json:"updateUserId"`                                     //修改人員
	Deleted      gorm.DeletedAt `gorm:"comment:軟刪除;type:datetime; json:"deleted"`                                        //軟刪除
	Picture      []Picture      `gorm:"many2many:carousel_picture;"`                                                     //圖片(多對多)
}
