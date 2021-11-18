package forestage

type CarouselDTO struct {
	Carousel []*CarouselData `json:"carousels"`
}

type CarouselData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Url    string `json:"url"`
	Weight int    `json:"weight"`
}
