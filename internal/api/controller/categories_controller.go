package controller

import (
	"componentmod/internal/services/api"

	"github.com/gin-gonic/gin"
)

var (
	CategoryList = Handler(GetCategory)
)

func GetCategory(c *gin.Context) (Data, error) {
	shService := api.GetShopeeService()
	return shService.GetCategoryList()
}
