package forestage

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	BaseForestageConfig = controller.Handler(GetConfig)
)

// @Summary Forestage config
// @Success 200 {json} json
// @Router /forestage/config [get]
func GetConfig(c *gin.Context) (controller.Data, error) {
	baseForestageService := forestage.GetBaseForestageService()
	return baseForestageService.GetBaseConfig()
}
