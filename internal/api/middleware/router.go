package middleware

import (
	"componentmod/internal/api/config"
	backstage "componentmod/internal/api/controller/backstagectl"
	forestages "componentmod/internal/api/controller/forestagectl"
	"componentmod/internal/api/middleware/validate"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {
	//----------------前台---------------

	//首頁輪播圖
	r.GET("/carousel/list", forestages.CarouselList)

	//相關設定檔
	r.GET("/forestage/config", forestages.BaseForestageConfig)

	//產品相關資料
	production := r.Group("/production")
	{
		//產品列表
		production.GET("/list", forestages.ProductionList)
		//產品詳細資料
		production.GET("/:id", forestages.ProductionById)
		//產品排名
		production.GET("/rank/:count", forestages.ProductionRank)
		//產品分類
		production.GET("/category/list", forestages.CategoryList)
	}

	//----------------後台---------------
	//顯示檔案或圖片
	r.GET("/file/:fileName", backstage.ShowFile)

	backstagePage := r.Group("/backstage")
	{
		//登入
		backstagePage.POST("/admin/login", backstage.BackstageLogin)

		//刷新 jwt Token
		backstagePage.POST("/jwt/refreshtoken", backstage.BackstageRefreshToken)

		//驗證 jwt Token
		backstagePage.GET("/jwt/check", backstage.BackstageCheckToken)

		//jwt 驗證通過
		backstagePage.Use(validate.JwtValidate())
		{
			//登出
			backstagePage.POST("/admin/logout", backstage.BackstageLogout)

			//菜單all權限列表
			backstagePage.GET("/role/all", backstage.RoleList)

			//菜單all權限列表
			backstagePage.GET("/menu/all", backstage.MenuTreeList)

			//菜單權限列表
			backstagePage.GET("/menu/list", backstage.MenuList)

			//拿取父類別 選項
			backstagePage.GET("/menu/parent/list", backstage.MenuParentList)

			//jwt 與 頁面權限 驗證通過
			user := backstagePage.Use()
			{
				user.Use(validate.AuthorityMenuValidateBYKey("user"))
				{
					// 使用者頁面
					user.GET("/user", backstage.UserShow)
					// 使用者 id
					user.GET("/user/:id", backstage.UserIndex)
				}

				user.Use(validate.AuthorityMenuValidateBYKey("user:create"))
				{
					// 使用者新增
					user.POST("/user/create", backstage.UserStore)
				}

				user.Use(validate.AuthorityMenuValidateBYKey("user:edit"))
				{
					// 使用者修改
					user.PUT("/user/edit/:id", backstage.UserUpdate)
				}

				user.Use(validate.AuthorityMenuValidateBYKey("user:delete"))
				{
					// 使用者刪除
					user.DELETE("/user/delete/:id", backstage.UserDestory)
				}

				user.Use(validate.AuthorityMenuValidateBYKey("user:password:edit"))
				{
					// 使用者密碼修改
					user.PUT("/user/password/edit/:id", backstage.UserPwdUpdate)
				}

				user.Use(validate.AuthorityMenuValidateBYKey("user:password:reset"))
				{
					// 使用者密碼重置
					user.PUT("/user/password/reset/:id", backstage.UserPwdSet)
				}
			}

			// 菜單
			menu := backstagePage.Use()
			{
				menu.Use(validate.AuthorityMenuValidateBYKey("menu"))
				{
					// 菜單頁面
					menu.GET("/menu", backstage.MenuShow)

					// 菜單 id
					menu.GET("/menu/:id", backstage.MenuIndex)
				}

				menu.Use(validate.AuthorityMenuValidateBYKey("menu:create"))
				{
					// // 菜單新增
					menu.POST("/menu/create", backstage.MenuStore)
				}
				menu.Use(validate.AuthorityMenuValidateBYKey("menu:edit"))
				{
					// // 菜單修改
					menu.PUT("/menu/edit/:id", backstage.MenuUpdate)
				}
				menu.Use(validate.AuthorityMenuValidateBYKey("menu:delete"))
				{
					// 菜單刪除
					menu.DELETE("/menu/delete/:id", backstage.MenuDestory)
				}
			}
			// 角色
			role := backstagePage.Use()
			{
				role.Use(validate.AuthorityMenuValidateBYKey("role"))
				{
					// 角色頁面
					role.GET("/role", backstage.RoleShow)

					// 角色 id
					role.GET("/role/:id", backstage.RoleIndex)
				}

				role.Use(validate.AuthorityMenuValidateBYKey("role:create"))
				{
					// 角色新增
					role.POST("/role/create", backstage.RoleStore)
				}
				role.Use(validate.AuthorityMenuValidateBYKey("role:edit"))
				{
					// 角色修改
					role.PUT("/role/edit/:id", backstage.RoleUpdate)
				}
				role.Use(validate.AuthorityMenuValidateBYKey("role:delete"))
				{
					// 角色刪除
					role.DELETE("/role/delete/:id", backstage.RoleDestory)
				}
			}

			// 清除cache
			cache := backstagePage.Use()
			{
				role.Use(validate.AuthorityMenuValidateBYKey("cache"))
				{
					// cache頁面
					cache.GET("/cache", backstage.CacheShow)
				}

				role.Use(validate.AuthorityMenuValidateBYKey("cache:delete"))
				{
					// cache特定刪除
					cache.DELETE("/cache/delete/:cacheName", backstage.CacheDestory)
				}

				role.Use(validate.AuthorityMenuValidateBYKey("cache:delete:any"))
				{
					// cache任意＊刪除
					cache.DELETE("/cache/any/delete/:cacheName", backstage.CacheAnyDestory)
				}
			}

			//輪播圖
			carousel := backstagePage.Use()
			{
				carousel.Use(validate.AuthorityMenuValidateBYKey("carousel"))
				{
					// 輪播圖頁面
					carousel.GET("/carousel", backstage.CarouselShow)
					carousel.GET("/carousel/:id", backstage.CarouselIndex)
				}

				carousel.Use(validate.AuthorityMenuValidateBYKey("carousel:create"))
				{
					// 角色新增
					carousel.POST("/carousel/create", backstage.CarouselStore)
				}
				carousel.Use(validate.AuthorityMenuValidateBYKey("carousel:edit"))
				{
					// 角色修改
					carousel.PUT("/carousel/edit/:id", backstage.CarouselUpdate)
				}
				carousel.Use(validate.AuthorityMenuValidateBYKey("carousel:delete"))
				{
					// 角色刪除
					carousel.DELETE("/carousel/delete/:id", backstage.CarouselDestory)
				}
			}

		}

	}

	//----------------swagger---------------
	// @title Gin Swagger Demo
	// @version 2.0
	// @description Swagger API.
	// @host localhost:80
	//swagger   http://localhost:80/swagger/index.html
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", config.WebPort)) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
