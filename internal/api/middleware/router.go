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

			//菜單all權限列表
			backstagePage.GET("/role/all", backstagectl.RoleList)

			//菜單all權限列表
			backstagePage.GET("/menu/all", backstagectl.MenuTreeList)

			//菜單權限列表
			backstagePage.GET("/menu/list", backstagectl.MenuList)

			//拿取父類別 選項
			backstagePage.GET("/menu/parent/list", backstagectl.MenuParentList)

			//jwt 與 頁面權限 驗證通過
			user := backstagePage.Use(validate.AuthorityMenuValidate())
			{
				// 使用者頁面
				user.GET("/user", backstagectl.UserShow)

				// 使用者 id
				user.GET("/user/:id", backstagectl.UserIndex)

				// 使用者新增
				user.POST("/user/create", backstagectl.UserStore)

				// 使用者修改
				user.PUT("/user/edit/:id", backstagectl.UserUpdate)

				// 使用者密碼修改
				user.PUT("/user/password/edit/:id", backstagectl.UserPwdUpdate)

				// 使用者密碼重置
				user.PUT("/user/password/reset/:id", backstagectl.UserPwdSet)

				// 使用者刪除
				user.DELETE("/user/delete/:id", backstagectl.UserDestory)
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
			// 角色
			role := backstagePage.Use(validate.AuthorityMenuValidate())
			{
				// 角色頁面
				role.GET("/role", backstagectl.RoleShow)

				// 角色 id
				role.GET("/role/:id", backstagectl.RoleIndex)

				// 角色新增
				role.POST("/role/create", backstagectl.RoleStore)

				// 角色修改
				role.PUT("/role/edit/:id", backstagectl.RoleUpdate)

				// 角色刪除
				role.DELETE("/role/delete/:id", backstagectl.RoleDestory)
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
