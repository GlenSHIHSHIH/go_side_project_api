package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/backstage"

	"github.com/gin-gonic/gin"
)

var (
	CacheDestory    = controller.Handler(CacheDelete)
	CacheAnyDestory = controller.Handler(CacheAnyDelete)
	CacheShow       = controller.Handler(Caches)
)

// @tags Forestage
// @Summary Cache Delete
// @accept application/json
// @Success 200
// @Router /backstage/cache/delete/{cacheName} [delete]
func CacheDelete(c *gin.Context) (controller.Data, error) {
	cacheName := c.Param("cacheName")
	cacheService := backstage.GetCacheService()
	return cacheService.DeleteCache(cacheName)
}

// @tags Forestage
// @Summary Cache Any Delete
// @accept application/json
// @Success 200
// @Router /backstage/cache/delete/{cacheName} [delete]
func CacheAnyDelete(c *gin.Context) (controller.Data, error) {
	cacheName := c.Param("cacheName")
	cacheService := backstage.GetCacheService()
	return cacheService.DeleteAnyCache(cacheName)
}

// @tags Forestage
// @Summary Cache View
// @accept application/json
// @Success 200
// @Router /backstage/cache [get]]
func Caches(c *gin.Context) (controller.Data, error) {
	cacheService := backstage.GetCacheService()
	return cacheService.GetCacheViewList()
}
