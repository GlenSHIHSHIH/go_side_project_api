package controller

import (
	"componentmod/internal/services/api"

	"github.com/gin-gonic/gin"
)

var (
	CategoriesList = Handler(GetCategories)
)

func GetCategories(c *gin.Context) (Data, error) {
	shService := api.GetShopeeService()
	return shService.Category()
}
