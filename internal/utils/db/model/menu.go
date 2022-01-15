package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"comment:菜單名稱;type:varchar(50)" json:"name"`                                       //菜單名稱
	Key          string         `gorm:"uniqueIndex;comment:菜單字符串;type:varchar(100)" json:"key"`                          //菜單字符串
	Url          string         `gorm:"comment:網址;type:varchar(300)" json:"url"`                                         //網址
	Feature      string         `gorm:"comment:功能(T=標題、P=頁面、F=按鍵功能);type:varchar(1)" json:"feature"`                     //功能(T=標題、P=頁面、F=按鍵功能)
	Weight       int            `gorm:"comment:權重(優先順序 重=高);type:int;default:0" json:"weight"`                           //權重
	Parent       int            `gorm:"comment:父類(id);type:int;default:0" json:"parent"`                                 //父類(id)
	Status       bool           `gorm:"comment:開關 (true=開啟);type:bool;default:true" json:"status"`                       //狀態(開關)
	CreateTime   time.Time      `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime   time.Time      `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
	CreateUserId int            `gorm:"comment:新增人員Id;type:int; json:"createUserId"`                                     //新增人員
	UpdateUserId int            `gorm:"comment:修改人員Id;type:int; json:"updateUserId"`                                     //修改人員
	Deleted      gorm.DeletedAt `gorm:"comment:軟刪除;type:datetime; json:"deleted"`                                        //軟刪除
	Remark       string         `gorm:"comment:備註;type:varchar(500);default:null" json:"remark"`                         //備註
	Role         []Role         `gorm:"many2many:role_menu;"`
}
