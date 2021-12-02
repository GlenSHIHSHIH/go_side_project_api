package forestagedto

type ProductionDetailDTO struct {
	Production *ProductionDetailData `json:"production"`
}

type ProductionDetailData struct {
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
	Attribute      string `json:"attribute"`
	LikedCount     int    `json:"likedCount"`
	HistoricalSold int    `json:"historicalSold"`
	Stock          int    `json:"stock"`
	CreateTime     string `json:"createTime"`
}
