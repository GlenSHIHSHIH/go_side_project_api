package router

import (
	"componentmod/internal/api/controller/forestage"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	production := r.Group("/production")
	{
		production.GET("/list", forestage.ProductionList)
		production.GET("/:id", forestage.ProductionById)
		production.GET("/rank/:count", forestage.ProductionRank)
	}

	carousel := r.Group("/carousel")
	{
		carousel.GET("/list", forestage.CarouselList)
	}

	r.GET("/category/list", forestage.CategoryList)

}
