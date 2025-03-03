package backstagedto

import "componentmod/internal/dto"

type MenuIdDTO struct {
	MenuById *MenuViewData `json:"menuById"`
}

type MenuCreateOrEditDTO struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
	Parent  string `json:"parent"`
	Weight  int    `json:"weight"`
	Status  bool   `json:"status"`
	Remark  string `json:"remark"`
}

type MenuViewData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
	Parent  string `json:"parent"`
	Weight  int    `json:"weight"`
	Status  bool   `json:"status"`
	Remark  string `json:"remark"`
}

type MenuViewListDTO struct {
	MenuViewList []*MenuViewData           `json:"menuViewList"`
	PageData     *dto.PageForMultSearchDTO `json:"pageData"`
}

type MenuData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
	Status  bool   `json:"status"`
	Parent  int    `json:"parent"`
}

type MenuDTO struct {
	Menu []*MenuNestData `json:"menu"`
}

type MenuNestData struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Key     string          `json:"key"`
	Url     string          `json:"url"`
	Feature string          `json:"feature"`
	Parent  int             `json:"parent"`
	Status  bool            `json:"status"`
	Child   []*MenuNestData `json:"child"`
}

// menu parent 後台下拉選單
type MenuParentDTO struct {
	MenuParentList []*MenuParentListDTO `json:"menuParentList"`
}

type MenuParentListDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
