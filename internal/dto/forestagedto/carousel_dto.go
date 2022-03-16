package forestagedto

import (
	"time"
)

type CarouselDTO struct {
	Carousel CarouselData   `json:"carousel"`
	Picture  []*PictureData `json:"picture"`
}

type CarouselData struct {
	Id           int       `json:"id"`
	CarouselName string    `json:"carouselName"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}

type CarouselPictureData struct {
	Id           int       `json:"id"`
	CarouselName string    `json:"carouselName"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	PictureName  string    `json:"pictureName"`
	Alt          string    `json:"alt"`
	Url          string    `json:"url"`
	Weight       int       `json:"weight"`
}
