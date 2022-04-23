package forestagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/services/api/forestage"

	"github.com/gin-gonic/gin"
)

var (
	CategoryList = controller.Handler(GetCategory)
)

// @tags Forestage
// @Summary Category
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} forestagedto.CategoryDTO
// @Router /category/list [get]
func GetCategory(c *gin.Context) (controller.Data, error) {
	categoryService := forestage.GetCategoryService()
	return categoryService.GetCategoryList()
}
