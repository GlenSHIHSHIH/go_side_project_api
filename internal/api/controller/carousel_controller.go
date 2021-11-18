package controller

import (
	"componentmod/internal/services/api"

	"github.com/gin-gonic/gin"
)

var (
	CarouselList = Handler(GetCarousel)
)

func GetCarousel(c *gin.Context) (Data, error) {
	shService := api.GetShopeeService()
	return shService.GetCarouselList()
}
