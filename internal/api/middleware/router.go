package middleware

import (
	"componentmod/internal/api/controller/backstagectl"
	"componentmod/internal/api/controller/forestagectl"
	"componentmod/internal/api/middleware/validate"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {

	//----------------前台---------------
	//產品相關資料
	production := r.Group("/production")
	{
		//產品列表
		production.GET("/list", forestagectl.ProductionList)
		//產品詳細資料
		production.GET("/:id", forestagectl.ProductionById)
		//產品排名
		production.GET("/rank/:count", forestagectl.ProductionRank)
		//產品分類
		production.GET("/category/list", forestagectl.CategoryList)
	}

	//首頁輪播圖
	r.GET("/carousel/list", forestagectl.CarouselList)

	//相關設定檔
	r.GET("/forestage/config", forestagectl.BaseForestageConfig)

	//----------------後台---------------
	backstagePage := r.Group("/backstage")
	{
		//登入
		backstagePage.POST("/admin/login", backstagectl.BackstageLogin)
		//刷新 jwt Token
		backstagePage.POST("/jwt/refreshtoken", backstagectl.BackstageRefreshToken)
		//驗證 jwt Token
		backstagePage.GET("/jwt/check", backstagectl.BackstageCheckToken)

		backstagePage.Use(validate.JwtValidate())
		{
			//登出
			backstagePage.POST("/admin/logout", backstagectl.BackstageLogout)
			//菜單權限列表
			backstagePage.GET("/menu/list", backstagectl.MenuList)
		}
		backstagePage.Use(validate.JwtValidate()).Use(validate.AuthorityMenuValidate())
		{
			//測試  尚未修正
			backstagePage.POST("/user/test", backstagectl.UserEdit)
			//新增使用者 (未詳細完成)
			backstagePage.POST("/user/create", backstagectl.UserCreate)
		}
	}

	//----------------swagger---------------
	// @title Gin Swagger Demo
	// @version 2.0
	// @description Swagger API.
	// @host localhost:80
	//swagger   http://localhost:80/swagger/index.html
	url := ginSwagger.URL("http://localhost:80/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
