package backstagedto

type UserDTO struct {
	Name      string `validate:"min=4" gorm:"comment:使用者名稱;type:varchar(30)" json:"name"`                 //使用者名稱
	LoginName string `validate:"min=4" gorm:"comment:登入帳號;type:varchar(30);uniqueIndex" json:"loginName"` //登入帳號
	Password  string `validate:"min=6" gorm:"comment:密碼;type:varchar(100)" json:"password"`               //密碼
	Email     string `gorm:"comment:Email;type:varchar(50)" json:"email"`                                 //Email
	Status    bool   `gorm:"comment:帳號狀態(false停用 true正常);type:bool;default:true" json:"status"`           //帳號狀態(false停用 true正常)
	UserType  bool   `gorm:"comment:是否為系統用戶;type:bool;default:true" json:"UserType"`                      //是否為系統用戶
	Remark    string `gorm:"comment:備註;type:varchar(500);default:null" json:"remark"`                     //備註
}
