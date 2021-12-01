package forestage

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	CarouselList = controller.Handler(GetCarousel)
)

// @Summary Carousel
// @Success 200 {json} json
// @Router /carousel/list [get]
func GetCarousel(c *gin.Context) (controller.Data, error) {
	carouselService := forestage.GetCarouselService()
	return carouselService.GetCarouselList()
}
