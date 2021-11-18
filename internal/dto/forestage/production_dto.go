package forestage

import "componentmod/internal/dto"

type ProductionDTO struct {
	ProductionList []*ProductionData `json:"productionList"`
	PageData       *dto.PageDTO      `json:"pageData"`
}

type ProductionByIdDTO struct {
	Production *ProductionData `json:"production"`
}

type ProductionData struct {
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
