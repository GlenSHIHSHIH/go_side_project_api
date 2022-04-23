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

// @tags Backstage-Menu
// @Summary Menu Create
// @accept application/json
// @Security BearerAuth
// @Success 200
// @Param json body backstagedto.MenuCreateOrEditDTO true "json"
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

// @tags Backstage-Menu
// @Summary Menu Edit
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Param json body backstagedto.MenuCreateOrEditDTO true "json"
// @Router /backstage/menu/edit/{id} [put]
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

// @tags Backstage-Menu
// @Summary Menu Delete
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Router /backstage/menu/delete/{id} [delete]
func MenuDelete(c *gin.Context) (controller.Data, error) {
	ids := strings.Split(c.Param("id"), ",")

	menuService := backstage.GetMenuService()
	return menuService.DeleteMenu(ids)
}

// @tags Backstage-Menu
// @Summary Menu By Id
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.MenuIdDTO
// @param id path int true "id"
// @Router /backstage/menu/{id} [get]
func MenuById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")

	menuService := backstage.GetMenuService()
	return menuService.GetMenuById(id)
}

// @tags Backstage-Menu
// @Summary Menu View
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.MenuViewListDTO
// @Param page query int true "int default" default(1)
// @Param pageLimit query int true "int enums" Enums(15,20,30,40,50)
// @Param sort query string true "string enums" Enums(asc,desc)
// @Param sortColumn query string true "string enums" Enums(id,key,url)
// @Param search query string false "string default" default()
// @Param searchCategory query string false "string default" default()
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

// @tags Backstage-Menu
// @Summary Menu List
// @accept application/json
// @Security BearerAuth
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

// @tags Backstage-Menu
// @Summary Menu Tree List
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.MenuDTO
// @Router /backstage/menu/parent/list [get]
func MenuTree(c *gin.Context) (controller.Data, error) {
	menuService := backstage.GetMenuService()
	return menuService.GetMenuAllList()
}

// @tags Backstage-Menu
// @Summary Menu Parent List
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.MenuDTO
// @Router /backstage/menu/parent/list [get]
func MenuParent(c *gin.Context) (controller.Data, error) {

	menuService := backstage.GetMenuService()
	return menuService.GetMenuParentList()
}
