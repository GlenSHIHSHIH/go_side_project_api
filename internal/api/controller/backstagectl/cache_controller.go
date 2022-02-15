package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/backstage"

	"github.com/gin-gonic/gin"
)

var (
	CacheDestory = controller.Handler(CacheDelete)
)

// @tags Forestage
// @Summary Cache Delete
// @accept application/json
// @Success 200
// @Router /cache/delete/{cacheName} [delete]
func CacheDelete(c *gin.Context) (controller.Data, error) {
	cacheName := c.Param("cacheName")
	cacheService := backstage.GetCacheService()
	return cacheService.DeleteCache(cacheName)
}
