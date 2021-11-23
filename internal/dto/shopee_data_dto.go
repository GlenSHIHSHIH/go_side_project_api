package dto

//要寫成 csv 的 production 資料結構
type ShopeeDataDTO struct {
	ProductId      int64
	Name           string
	Description    string
	Options        []Options
	Categories     string
	Image          string
	Images         string
	Url            string
	Price          int64
	PriceMin       int64
	Attribute      string
	LikedCount     int64
	HistoricalSold int64
	Stock          int64
}

type Options struct {
	Name   string
	Option []string
}
