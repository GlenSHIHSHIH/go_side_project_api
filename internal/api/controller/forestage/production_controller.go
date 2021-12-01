package forestage

import (
	"componentmod/internal/api/controller"
	errorCode "componentmod/internal/api/error_code"
	"componentmod/internal/dto"
	"componentmod/internal/services/api/forestage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ProductionList = controller.Handler(GetProduction)
	ProductionById = controller.Handler(GetProductionById)
)

// @Summary Production list
// @Success 200 {json} json
// @Param page query int true "int default" default(1)
// @Param pageLimit query int true "int enums" Enums(15,20,30,40,50)
// @Param sort query string true "string enums" Enums(asc,desc)
// @Param sortColumn query string true "string enums" Enums(PName,PId,PCategory,PCreTime)
// @Param search query string false "string default" default()
// @Param searchCategory query string false "string default" default()
// @Router /production/list [get]
func GetProduction(c *gin.Context) (controller.Data, error) {
	var shopeePageDTO = &dto.PageDTO{
		Page:           1,
		PageLimit:      20,
		Sort:           "asc",
		SortColumn:     "PId",
		Search:         "",
		SearchCategory: "",
	}
	err := c.Bind(shopeePageDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorCode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	productionService := forestage.GetProductionService()
	return productionService.GetProductionList(shopeePageDTO)
}

// @Summary Production detail
// @Id 1
// @Success 200 {json} json
// @param id path int true "id"
// @Router /production/{id} [get]
func GetProductionById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")
	productionService := forestage.GetProductionService()
	return productionService.GetProductionById(id)
}
