package forestage

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	BaseForestageConfig = controller.Handler(GetConfig)
)

func GetConfig(c *gin.Context) (controller.Data, error) {
	baseForestageService := forestage.GetBaseForestageService()
	return baseForestageService.GetBaseConfig()
}
