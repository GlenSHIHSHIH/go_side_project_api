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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	MenuList       = controller.Handler(GetMenu)
	MenuParentList = controller.Handler(GetMenuParentList)
	Menu           = controller.Handler(Menus)
	MenuId         = controller.Handler(MenuById)
	MenuDestory    = controller.Handler(MenuDelete)
)

// @tags Backstage
// @Summary Menu Delete
// @accept application/json
// @Success 200 {object}
// @Router /backstage/menu/delete [delete]
func MenuDelete(c *gin.Context) (controller.Data, error) {
	ids := strings.Split(c.Param("id"), ",")

	menuService := backstage.GetMenuService()
	return menuService.DeleteMenu(ids)
}

// @tags Backstage
// @Summary Menu By Id
// @accept application/json
// @Success 200 {object} backstagedto.MenuIdDTO
// @Router /backstage/menu/id [get]
func MenuById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")

	menuService := backstage.GetMenuService()
	return menuService.GetMenuById(id)
}

// @tags Backstage
// @Summary Menu View
// @accept application/json
// @Success 200 {object} backstagedto.MenuViewListDTO
// @Router /backstage/menu [get]
func Menus(c *gin.Context) (controller.Data, error) {
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
// @Summary Menu Parent List
// @accept application/json
// @Success 200 {object} backstagedto.MenuDTO
// @Router /backstage/menu/parent/list [get]
func GetMenuParentList(c *gin.Context) (controller.Data, error) {

	menuService := backstage.GetMenuService()
	return menuService.GetMenuParentList()
}
