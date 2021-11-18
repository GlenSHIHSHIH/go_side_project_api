package router

import (
	"componentmod/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	production := r.Group("/production")
	{

		production.GET("/list", controller.ProductionList)
		production.GET("/:id", controller.ProductionById)
	}

	carousel := r.Group("/carousel")
	{
		carousel.GET("/list", controller.CarouselList)
	}

	r.GET("/categories/list", controller.CategoryList)

}
