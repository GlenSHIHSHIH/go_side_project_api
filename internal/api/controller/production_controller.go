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
	CategoriesList = Handler(GetCategories)
)

func GetProduction(c *gin.Context) (Data, error) {
	var shopeeProduction = &dto.ShopeeProductionInDTO{
		Page:           1,
		PageLimit:      20,
		Sort:           "asc",
		SortColumn:     "PId",
		Search:         "",
		SearchCategory: "",
	}
	err := c.Bind(&shopeeProduction)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorCode.PARAMETER_ERROR)
		log.Warn(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorCode.PARAMETER_ERROR_CODE, errorCode.PARAMETER_ERROR)
	}

	shService := api.GetShopeeService()
	return shService.Production(shopeeProduction)
}

func GetCategories(c *gin.Context) (Data, error) {
	shService := api.GetShopeeService()
	return shService.Category()
}
