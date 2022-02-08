package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	ShowFile = controller.WebHandler(GetFile)
)

// @tags Forestage
// @Summary Get File (image...)
// @accept application/json
// @Success 200
// @Router /file/{id} [get]
func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	fileService := forestage.GetFileService()
	fileService.GetFile(fileName, c)
}
