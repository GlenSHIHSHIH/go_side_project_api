package dto

type ShopeeProductionInDTO struct {
	Page           int    `form:"page"`
	PageLimit      int    `form:"pageLimit"`
	Sort           string `form:"sort"`
	SortColumn     string `form:"sortColumn"`
	Search         string `form:"search"`
	SearchCategory string `form:"searchCategory"`
	// Filter       PFilter `form:"filter"`
}

// type PFilter struct {
// 	Column string `form:"column"`
// 	Value  string `form:"value"`
// }

type ShopeeProductionOutDTO struct {
	ProductId   int64  `json:"productId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Options     string `json:"options"`
	Categories  string `json:"categories"`
	Image       string `json:"image"`
	Images      string `json:"images"`
	Url         string `json:"url"`
	CreateTime  string `json:"createTime"`
}
