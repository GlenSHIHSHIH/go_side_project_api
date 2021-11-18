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
		log.Warn(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	productionService := forestage.GetProductionService()
	return productionService.GetProductionList(shopeePageDTO)
}

func GetProductionById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")
	productionService := forestage.GetProductionService()
	return productionService.GetProductionById(id)
}
