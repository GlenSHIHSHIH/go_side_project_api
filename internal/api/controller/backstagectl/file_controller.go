package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/backstage"

	"github.com/gin-gonic/gin"
)

var (
	ShowFile = controller.WebHandler(GetFile)
)

// @tags Forestage
// @Summary Get File (image...)
// @accept application/json
// @Security BearerAuth
// @Success 200
// @Router /file/{id} [get]
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	fileService := backstage.GetFileService()
	fileService.GetFile(fileName, c)
}
