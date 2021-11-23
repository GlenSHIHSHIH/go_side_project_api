package forestage

type ProductionRankDTO struct {
	ProductionList []*ProductionRankData `json:"productionList"`
}

type ProductionRankData struct {
	Id             int    `json:"id"`
	ProductId      int64  `json:"productId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Options        string `json:"options"`
	Categories     string `json:"categories"`
	Image          string `json:"image"`
	Images         string `json:"images"`
	Url            string `json:"url"`
	Price          int64  `json:"price"`
	PriceMin       int64  `json:"priceMin"`
	CreateTime     string `json:"createTime"`
	Weight         int    `json:"weight"`
	Amount         int    `json:"amount"`
	LikedCount     int    `json:"likedCount"`
	HistoricalSold int    `json:"historicalSold"`
	Stock          int    `json:"stock"`
}
