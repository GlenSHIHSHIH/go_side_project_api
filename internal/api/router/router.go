package router

import (
	"componentmod/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	root := r.Group("/")
	{

		root.Any("", controller.ProductionList)
		// root.GET("", controller.ProductionList)
	}
	// root.Use(middleware){	}

}
