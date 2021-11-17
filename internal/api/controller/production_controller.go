package controller

import (
	errorCode "componentmod/internal/api/error_code"
	"componentmod/internal/dto"
	"componentmod/internal/services/api"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ProductionList = Handler(GetProduction)
	ProductionById = Handler(GetProductionById)
)

func GetProduction(c *gin.Context) (Data, error) {
	var shopeePageDTO = &dto.ShopeePageDTO{
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

	shService := api.GetShopeeService()
	return shService.Production(shopeePageDTO)
}

func GetProductionById(c *gin.Context) (Data, error) {
	// err := c.Param("name")

	return nil, nil
}
