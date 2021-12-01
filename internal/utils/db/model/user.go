package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id            int            `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"comment:使用者名稱;type:varchar(30)" json:"name"`                                      //使用者名稱
	LoginName     string         `gorm:"comment:登入帳號;type:varchar(30)" json:"loginName"`                                  //登入帳號
	Password      string         `gorm:"comment:密碼;type:varchar(50)" json:"password"`                                     //密碼
	Email         string         `gorm:"comment:Email;type:varchar(50)" json:"email"`                                     //Email
	Status        bool           `gorm:"comment:帳號狀態(false停用 true正常);type:bool;default:true" json:"status"`               //帳號狀態(false停用 true正常)
	UserType      bool           `gorm:"comment:是否為系統用戶;type:bool;default:true" json:"UserType"`                          //是否為系統用戶
	LoginIP       string         `gorm:"comment:ip;type:varchar(20)" json:"loginIP"`                                      //loginIP
	PwdUpdateTime time.Time      `gorm:"comment:密碼最後更新時間;type:datetime;default:CURRENT_TIMESTAMP" json:"pwdUpdateTime"`   //密碼最後更新時間
	LoginTime     time.Time      `gorm:"comment:最後登入時間;type:datetime;default:CURRENT_TIMESTAMP" json:"loginTime"`         //最後登入時間
	CreateTime    time.Time      `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime    time.Time      `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
	CreateUserId  int            `gorm:"comment:新增人員Id;type:int" json:"createUserId"`                                     //新增人員
	UpdateUserId  int            `gorm:"comment:修改人員Id;type:int" json:"updateUserId"`                                     //修改人員
	Deleted       gorm.DeletedAt `gorm:"comment:軟刪除;type:datetime;default:null" json:"deleted"`                           //軟刪除
	Remark        string         `gorm:"comment:備註;type:varchar(500);default:null" json:"remark"`                         //備註
	Role          []Role         `gorm:"many2many:user_roles;"`                                                           //角色(多對多)
}
