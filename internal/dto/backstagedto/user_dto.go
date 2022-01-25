package backstagedto

import (
	"componentmod/internal/dto"
	"time"
)

type UserDTO struct {
	Name      string `validate:"min=4" json:"name"`      //使用者名稱
	LoginName string `validate:"min=4" json:"loginName"` //登入帳號
	Password  string `validate:"min=6" json:"password"`  //密碼
	Email     string `json:"email"`                      //Email
	Status    bool   `json:"status"`                     //帳號狀態(false停用 true正常)
	UserType  bool   `json:"UserType"`                   //是否為系統用戶
	Remark    string `json:"remark"`                     //備註
}

type UserIdDTO struct {
	UserById *UserCreateOrEditDTO `json:"userById"`
}

type UserCreateOrEditDTO struct {
	Id        int    `json:"id"`
	Name      string `validate:"min=4" json:"name"`      //使用者名稱
	LoginName string `validate:"min=4" json:"loginName"` //登入帳號
	Password  string `validate:"min=6" json:"password"`  //密碼
	Email     string `json:"email"`                      //Email
	Status    bool   `json:"status"`                     //帳號狀態(false停用 true正常)
	UserType  bool   `json:"UserType"`                   //是否為系統用戶
	Remark    string `json:"remark"`                     //備註
	Select    []int  `json:"select"`                     //選的role
}

type UserViewData struct {
	Id         int       `json:"id"`
	Name       string    `validate:"min=4" json:"name"`      //使用者名稱
	LoginName  string    `validate:"min=4" json:"loginName"` //登入帳號
	Password   string    `validate:"min=6" json:"password"`  //密碼
	Email      string    `json:"email"`                      //Email
	Status     bool      `json:"status"`                     //帳號狀態(false停用 true正常)
	UserType   bool      `json:"UserType"`                   //是否為系統用戶
	Remark     string    `json:"remark"`                     //備註
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser string    `json:"createUser"`
	UpdateUser string    `json:"updateUser"`
}

type UserListDTO struct {
	UserList []*UserViewData           `json:"roleList"`
	PageData *dto.PageForMultSearchDTO `json:"pageData"`
}
