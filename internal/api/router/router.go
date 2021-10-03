package router

import (
	"componentmod/internal/api/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	root := r.Group("/")
	{
		root.GET("*", controller.Handler(controller.Production))
	}
	// root.Use(middleware){	}

}
