package backstagedto

import (
	"componentmod/internal/dto"
	"time"
)

type RoleIdDTO struct {
	RoleById *RoleCreateOrEditDTO `json:"roleById"`
}

type RoleCreateOrEditDTO struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Weight int    `json:"weight"`
	Status bool   `json:"status"`
	Remark string `json:"remark"`
	Select []int  `json:"select"`
}

type RoleViewData struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Key        string    `json:"key"`
	Weight     int       `json:"weight"`
	Status     bool      `json:"status"`
	Remark     string    `json:"remark"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser string    `json:"createUser"`
	UpdateUser string    `json:"updateUser"`
}

type RoleListDTO struct {
	RoleList []*RoleViewData           `json:"roleList"`
	PageData *dto.PageForMultSearchDTO `json:"pageData"`
}
