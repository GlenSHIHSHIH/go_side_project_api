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
	RoleList    = controller.Handler(RolesList)
	RoleShow    = controller.Handler(Roles)
	RoleIndex   = controller.Handler(RoleById)
	RoleDestory = controller.Handler(RoleDelete)
	RoleStore   = controller.Handler(RoleCreate)
	RoleUpdate  = controller.Handler(RoleEdit)
)

//table list

// @tags Backstage-Role
// @Summary Role View
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.RoleListDTO
// @Param page query int true "int default" default(1)
// @Param pageLimit query int true "int enums" Enums(15,20,30,40,50)
// @Param sort query string true "string enums" Enums(asc,desc)
// @Param sortColumn query string true "string enums" Enums(id,key)
// @Param search query string false "string default" default()
// @Param searchCategory query string false "string default" default()
// @Router /backstage/role [get]
func Roles(c *gin.Context) (controller.Data, error) {
	search := c.QueryMap("search")
	var pageForMultSearchDTO = GetPageMultSearchDefaultDTO()
	pageForMultSearchDTO.Search = search

	err := c.Bind(pageForMultSearchDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	roleService := backstage.GetRoleService()
	return roleService.GetRoleViewList(pageForMultSearchDTO)
}

// @tags Backstage-Role
// @Summary Role List
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.RoleIdDTO
// @param id path int true "id"
// @Router /backstage/role/{id} [get]
func RolesList(c *gin.Context) (controller.Data, error) {

	roleService := backstage.GetRoleService()
	return roleService.GetRoleList()
}

// @tags Backstage-Role
// @Summary Role By Id
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.RoleIdDTO
// @param id path int true "id"
// @Router /backstage/role/{id} [get]
func RoleById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")

	roleService := backstage.GetRoleService()
	return roleService.GetRoleById(id)
}

// @tags Backstage-Role
// @Summary Role Delete
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Router /backstage/role/delete/{id} [delete]
func RoleDelete(c *gin.Context) (controller.Data, error) {
	ids := strings.Split(c.Param("id"), ",")

	roleService := backstage.GetRoleService()
	return roleService.DeleteRole(ids)
}

// @tags Backstage-Role
// @Summary Role Create
// @accept application/json
// @Security BearerAuth
// @Success 200
// @Param json body backstagedto.RoleCreateOrEditDTO true "json"
// @Router /backstage/role/create [post]
func RoleCreate(c *gin.Context) (controller.Data, error) {

	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	var roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO
	err = c.Bind(&roleCreateOrEditDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	roleService := backstage.GetRoleService()
	return roleService.CreateRole(userInfo, roleCreateOrEditDTO)
}

// @tags Backstage-Role
// @Summary Role Edit
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Param json body backstagedto.RoleCreateOrEditDTO true "json"
// @Router /backstage/role/edit/{id} [put]
func RoleEdit(c *gin.Context) (controller.Data, error) {
	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	id := c.Param("id")

	var roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO
	err = c.Bind(&roleCreateOrEditDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}
	roleService := backstage.GetRoleService()
	return roleService.EditRole(userInfo, id, roleCreateOrEditDTO)
}
