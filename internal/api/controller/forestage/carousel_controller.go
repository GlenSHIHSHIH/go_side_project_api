package forestage

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	CarouselList = controller.Handler(GetCarousel)
)

func GetCarousel(c *gin.Context) (controller.Data, error) {
	carouselService := forestage.GetCarouselService()
	return carouselService.GetCarouselList()
}
