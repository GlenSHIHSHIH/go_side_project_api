package backstagedto

import (
	"componentmod/internal/dto"
	"time"
)

type UserEditPwdDTO struct {
	Id          int    `json:"id"`
	Type        int    `json:"type"`                         //1.重置密碼 2.修改密碼
	OrgPassword string `json:"orgPassword"`                  //原始密碼
	Password    string `validate:"min=6" json:"newPassword"` //新密碼
}

type UserIdDTO struct {
	UserById *UserCreateOrEditDTO `json:"userById"`
}

type UserCreateOrEditDTO struct {
	Id        int      `json:"id"`
	Name      string   `validate:"min=4" json:"name"`      //使用者名稱
	LoginName string   `validate:"min=4" json:"loginName"` //登入帳號
	Password  string   `validate:"min=6" json:"password"`  //密碼
	Email     string   `json:"email"`                      //Email
	Status    bool     `json:"status"`                     //帳號狀態(false停用 true正常)
	UserType  bool     `json:"userType"`                   //是否為系統用戶
	Remark    string   `json:"remark"`                     //備註
	Select    []string `json:"select"`                     //選的role
}

type UserViewData struct {
	Id         int       `json:"id"`
	Name       string    `validate:"min=4" json:"name"`      //使用者名稱
	LoginName  string    `validate:"min=4" json:"loginName"` //登入帳號
	Email      string    `json:"email"`                      //Email
	Status     bool      `json:"status"`                     //帳號狀態(false停用 true正常)
	UserType   bool      `json:"userType"`                   //是否為系統用戶
	Remark     string    `json:"remark"`                     //備註
	Role       string    `json:"role"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser string    `json:"createUser"`
	UpdateUser string    `json:"updateUser"`
}

type UserListDTO struct {
	UserList []*UserViewData           `json:"userList"`
	PageData *dto.PageForMultSearchDTO `json:"pageData"`
}
