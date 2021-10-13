package router

import (
	"componentmod/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	root := r.Group("/")
	{

		root.GET("production/list", controller.ProductionList)
		root.GET("categoriesList/list", controller.CategoriesList)
		// root.GET("", controller.ProductionList)
	}
	// root.Use(middleware){	}

}
