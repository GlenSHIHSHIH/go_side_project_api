package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/middleware/validate"
	"componentmod/internal/services/api/backstage"

	"github.com/gin-gonic/gin"
)

var (
	MenuList = controller.Handler(GetMenu)
)

// // @tags Backstage
// // @Summary Menu
// // @accept application/json
// // @Success 200 {object} Backstage.CarouselDTO
// // @Router /carousel/list [get]
func GetMenu(c *gin.Context) (controller.Data, error) {

	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}
	menuService := backstage.GetMenuService()
	return menuService.GetMenuListByUserId(userInfo.Id)
}
