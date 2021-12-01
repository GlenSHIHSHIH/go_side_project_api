package router

import (
	"componentmod/internal/api/controller/forestage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.GET("/forestage/config", forestage.BaseForestageConfig)

	// @title Gin Swagger Demo
	// @version 2.0
	// @description Swagger API.
	// @host localhost:80
	//swagger   http://localhost:80/swagger/index.html
	url := ginSwagger.URL("http://localhost:80/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
