package forestage

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	CategoryList = controller.Handler(GetCategory)
)

func GetCategory(c *gin.Context) (controller.Data, error) {
	categoryService := forestage.GetCategoryService()
	return categoryService.GetCategoryList()
}
