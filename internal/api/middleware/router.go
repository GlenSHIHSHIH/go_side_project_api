package middleware

import (
	"componentmod/internal/api/controller/backstagectl"
	"componentmod/internal/api/controller/forestagectl"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {
	production := r.Group("/production")
	{
		production.GET("/list", forestagectl.ProductionList)
		production.GET("/:id", forestagectl.ProductionById)
		production.GET("/rank/:count", forestagectl.ProductionRank)
	}

	carousel := r.Group("/carousel")
	{
		carousel.GET("/list", forestagectl.CarouselList)
	}

	r.GET("/category/list", forestagectl.CategoryList)
	r.GET("/forestage/config", forestagectl.BaseForestageConfig)

	//登入 / 登出
	r.POST("/backstage/admin/login", backstagectl.BackstageLogin)
	r.POST("/backstage//admin/logout", backstagectl.BackstageLogout)

	r.POST("/backstage/jwt/refreshtoken", backstagectl.BackstageRefreshToken)

	//後台
	backstagePage := r.Group("/backstage")
	{
		backstagePage.Use(authorityJwtMenuCheck())
		{
			backstagePage.POST("/user/test", backstagectl.UserEdit)
		}
		// backstagePage.Use(authorityJwtMenuCheck()){
		backstagePage.POST("/user/create", backstagectl.UserCreate)
		// }

		// r.GET("/backstage/login", backstage.UserLogin)
	}

	// @title Gin Swagger Demo
	// @version 2.0
	// @description Swagger API.
	// @host localhost:80
	//swagger   http://localhost:80/swagger/index.html
	url := ginSwagger.URL("http://localhost:80/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
