package forestagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
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

// @tags Forestage
// @Summary Production list
// @accept application/json
// @Success 200 {object} forestagedto.ProductionDTO
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
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	productionService := forestage.GetProductionService()
	return productionService.GetProductionList(shopeePageDTO)
}

// @tags Forestage
// @Summary Production detail
// @accept application/json
// @Id 1
// @Success 200 {object} forestagedto.ProductionDetailDTO
// @param id path int true "id"
// @Router /production/{id} [get]
func GetProductionById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")
	productionService := forestage.GetProductionService()
	return productionService.GetProductionById(id)
}
