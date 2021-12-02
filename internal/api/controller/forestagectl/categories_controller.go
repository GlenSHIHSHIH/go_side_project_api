package forestagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	CategoryList = controller.Handler(GetCategory)
)

// @Summary Category
// @Success 200 {json} json
// @Router /category/list [get]
func GetCategory(c *gin.Context) (controller.Data, error) {
	categoryService := forestage.GetCategoryService()
	return categoryService.GetCategoryList()
}
