package dto

type ShopeeCarouselDTO struct {
	Carousel []*ShopeeCarouselData `json:"carousels"`
}

type ShopeeCarouselData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Url    string `json:"url"`
	Weight int    `json:"weight"`
}
