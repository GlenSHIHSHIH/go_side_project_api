package backstagedto

import "componentmod/internal/dto"

type MenuViewDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
	Parent  string `json:"parent"`
	Weight  int    `json:"weight"`
	Status  bool   `json:"status"`
}

type MenuViewListDTO struct {
	MenuViewList []*MenuViewDTO            `json:"menuViewList"`
	PageData     *dto.PageForMultSearchDTO `json:"pageData"`
}

type MenuData struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
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
	Child   []*MenuNestData `json:"child"`
}
