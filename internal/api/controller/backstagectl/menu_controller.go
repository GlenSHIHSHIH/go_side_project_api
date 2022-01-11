package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/api/middleware/validate"
	"componentmod/internal/dto"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	MenuList = controller.Handler(GetMenu)
	MenuView = controller.Handler(GetMenuView)
)

// @tags Backstage
// @Summary Menu List
// @accept application/json
// @Success 200 {object} backstagedto.MenuDTO
// @Router /backstage/menu/list [get]
func GetMenu(c *gin.Context) (controller.Data, error) {

	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}
	menuService := backstage.GetMenuService()
	return menuService.GetMenuNestList(userInfo.Id)
}

// @tags Backstage
// @Summary Menu View
// @accept application/json
// @Success 200 {object} backstagedto.MenuViewListDTO
// @Router /backstage/menu [get]
func GetMenuView(c *gin.Context) (controller.Data, error) {
	search := c.QueryMap("search")
	var pageForMultSearchDTO = &dto.PageForMultSearchDTO{
		Page:       1,
		PageLimit:  20,
		Sort:       "asc",
		SortColumn: "id",
		Search:     search,
	}

	err := c.Bind(pageForMultSearchDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	menuService := backstage.GetMenuService()
	return menuService.GetMenuViewList(pageForMultSearchDTO)
}
