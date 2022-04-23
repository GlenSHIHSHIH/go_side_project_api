package forestagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	BaseForestageConfig = controller.Handler(GetConfig)
)

// @tags Forestage
// @Summary Forestage config
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} forestagedto.BaseForestageConfigDTO
// @Router /forestage/config [get]
func GetConfig(c *gin.Context) (controller.Data, error) {
	baseForestageService := forestage.GetBaseForestageService()
	return baseForestageService.GetBaseConfig()
}
