package dto

type ShopeeProductionDTO struct {
	ProductionList []*ShopeeProductionData `json:"productionList"`
	PageData       *ShopeePageDTO          `json:"pageData"`
}

type ShopeeProductionByIdDTO struct {
	Production *ShopeeProductionData `json:"production"`
}

type ShopeeProductionData struct {
	ProductId   int64  `json:"productId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Options     string `json:"options"`
	Categories  string `json:"categories"`
	Image       string `json:"image"`
	Images      string `json:"images"`
	Url         string `json:"url"`
	Price       int64  `json:"price"`
	PriceMin    int64  `json:"priceMin"`
	CreateTime  string `json:"createTime"`
}
