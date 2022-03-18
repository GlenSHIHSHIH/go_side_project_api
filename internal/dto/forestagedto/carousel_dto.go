package forestagedto

import (
	"time"
)

type CarouselDTO struct {
	Carousel *CarouselData  `json:"carousel"`
	Picture  []*PictureData `json:"picture"`
}

type CarouselData struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
