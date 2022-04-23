package forestagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	CarouselList = controller.Handler(GetCarousel)
)

// @tags Forestage
// @Summary Carousel
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} forestagedto.CarouselDTO
// @Router /carousel/list [get]
func GetCarousel(c *gin.Context) (controller.Data, error) {
	carouselService := forestage.GetCarouselService()
	return carouselService.GetCarouselList()
}
