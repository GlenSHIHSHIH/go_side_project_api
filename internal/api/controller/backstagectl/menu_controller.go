package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/api/middleware/validate"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	MenuTreeList   = controller.Handler(MenuTree)
	MenuList       = controller.Handler(Menu)
	MenuParentList = controller.Handler(MenuParent)
	MenuShow       = controller.Handler(Menus)
	MenuIndex      = controller.Handler(MenuById)
	MenuDestory    = controller.Handler(MenuDelete)
	MenuStore      = controller.Handler(MenuCreate)
	MenuUpdate     = controller.Handler(MenuEdit)
)

// @tags Backstage
// @Summary Menu Create
// @accept application/json
// @Success 200 {object}
// @Router /backstage/menu/create [post]
func MenuCreate(c *gin.Context) (controller.Data, error) {

	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	var menuCreateOrEditDTO *backstagedto.MenuCreateOrEditDTO
	err = c.Bind(&menuCreateOrEditDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	menuService := backstage.GetMenuService()
	return menuService.CreateMenu(userInfo, menuCreateOrEditDTO)
}

// @tags Backstage
// @Summary Menu Edit
// @accept application/json
// @Success 200 {object}
// @Router /backstage/menu/edit [put]
func MenuEdit(c *gin.Context) (controller.Data, error) {
	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	id := c.Param("id")

	var menuCreateOrEditDTO *backstagedto.MenuCreateOrEditDTO
	err = c.Bind(&menuCreateOrEditDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}
	menuService := backstage.GetMenuService()
	return menuService.EditMenu(userInfo, id, menuCreateOrEditDTO)
}

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
	var pageForMultSearchDTO = GetPageMultSearchDefaultDTO()
	pageForMultSearchDTO.Search = search

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
func Menu(c *gin.Context) (controller.Data, error) {

	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}
	menuService := backstage.GetMenuService()
	return menuService.GetMenuNestList(userInfo.Id)
}

// @tags Backstage
// @Summary Menu Tree List
// @accept application/json
// @Success 200 {object} backstagedto.MenuDTO
// @Router /backstage/menu/parent/list [get]
func MenuTree(c *gin.Context) (controller.Data, error) {
	menuService := backstage.GetMenuService()
	return menuService.GetMenuNestList(userInfo.Id)
}

// @tags Backstage
// @Summary Menu Parent List
// @accept application/json
// @Success 200 {object} backstagedto.MenuDTO
// @Router /backstage/menu/parent/list [get]
func MenuParent(c *gin.Context) (controller.Data, error) {

	menuService := backstage.GetMenuService()
	return menuService.GetMenuParentList()
}
