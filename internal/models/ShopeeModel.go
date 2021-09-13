package models

//要寫成 csv 的 production 資料結構
type ShopeeDataModel struct {
	ProductId   int64
	Name        string
	Description string
	Option      []Options
	Categories  string
	Image       string
	Images      string
	Url         string
}

type Options struct {
	Name   string
	Option []string
}
