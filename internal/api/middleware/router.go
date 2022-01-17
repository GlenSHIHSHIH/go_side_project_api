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

	//首頁輪播圖
	r.GET("/carousel/list", forestagectl.CarouselList)

	//相關設定檔
	r.GET("/forestage/config", forestagectl.BaseForestageConfig)

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

	//----------------後台---------------
	backstagePage := r.Group("/backstage")
	{
		//登入
		backstagePage.POST("/admin/login", backstagectl.BackstageLogin)

		//刷新 jwt Token
		backstagePage.POST("/jwt/refreshtoken", backstagectl.BackstageRefreshToken)

		//驗證 jwt Token
		backstagePage.GET("/jwt/check", backstagectl.BackstageCheckToken)

		//jwt 驗證通過
		backstagePage.Use(validate.JwtValidate())
		{
			//登出
			backstagePage.POST("/admin/logout", backstagectl.BackstageLogout)
			//菜單權限列表
			backstagePage.GET("/menu/list", backstagectl.MenuList)
			//拿取父類別 選項
			backstagePage.GET("/menu/parent/list", backstagectl.MenuParentList)

			//jwt 與 頁面權限 驗證通過
			user := backstagePage.Use(validate.AuthorityMenuValidate())
			{

				//user
				//測試  尚未修正
				user.POST("/user/delete", backstagectl.UserEdit)
				//新增使用者 (未詳細完成)
				user.POST("/user/create", backstagectl.UserCreate)
			}

			// 菜單
			menu := backstagePage.Use(validate.AuthorityMenuValidate())
			{
				// 菜單頁面
				menu.GET("/menu", backstagectl.MenuShow)

				// 菜單 id
				menu.GET("/menu/:id", backstagectl.MenuIndex)

				// // 菜單新增
				menu.POST("/menu/create", backstagectl.MenuStore)

				// // 菜單修改
				menu.PUT("/menu/edit/:id", backstagectl.MenuUpdate)

				// 菜單刪除
				menu.DELETE("/menu/delete/:id", backstagectl.MenuDestory)
			}

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
