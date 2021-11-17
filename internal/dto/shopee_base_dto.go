package dto

type ShopeePageDTO struct {
	Page           int    `form:"page" json:"page"`
	PageLimit      int    `form:"pageLimit" json:"pageLimit"`
	Count          int64  `form:"count" json:"count"`
	Sort           string `form:"sort" json:"sort"`
	SortColumn     string `form:"sortColumn" json:"sortColumn"`
	Search         string `form:"search" json:"search"`
	SearchCategory string `form:"searchCategory" json:"searchCategory"`
}
