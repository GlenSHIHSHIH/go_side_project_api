package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	RoleShow    = controller.Handler(Roles)
	RoleIndex   = controller.Handler(RoleById)
	RoleDestory = controller.Handler(RoleDelete)
	// RoleStore   = controller.Handler(RoleCreate)
	// RoleUpdate  = controller.Handler(RoleEdit)
)

//table list

// @tags Backstage
// @Summary Role View
// @accept application/json
// @Success 200 {object} backstagedto.RoleListDTO
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

//find by id

// @tags Backstage
// @Summary Role By Id
// @accept application/json
// @Success 200 {object} backstagedto.RoleIdDTO
// @Router /backstage/role/id [get]
func RoleById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")

	roleService := backstage.GetRoleService()
	return roleService.GetRoleById(id)
}

//delete by id

// @tags Backstage
// @Summary Role Delete
// @accept application/json
// @Success 200 {object}
// @Router /backstage/role/delete [delete]
func RoleDelete(c *gin.Context) (controller.Data, error) {
	ids := strings.Split(c.Param("id"), ",")

	roleService := backstage.GetRoleService()
	return roleService.DeleteRole(ids)
}

// // @tags Backstage
// // @Summary Role Create
// // @accept application/json
// // @Success 200 {object}
// // @Router /backstage/role/create [post]
// func RoleCreate(c *gin.Context) (controller.Data, error) {

// 	userInfo, err := validate.UserInfoValidate(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO
// 	err = c.Bind(&roleCreateOrEditDTO)
// 	if err != nil {
// 		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
// 		log.Error(fmt.Sprintf("%+v", errData))
// 		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
// 	}

// 	roleService := backstage.GetRoleService()
// 	return roleService.CreateRole(userInfo, roleCreateOrEditDTO)
// }

// // @tags Backstage
// // @Summary Role Edit
// // @accept application/json
// // @Success 200 {object}
// // @Router /backstage/role/edit [put]
// func RoleEdit(c *gin.Context) (controller.Data, error) {
// 	userInfo, err := validate.UserInfoValidate(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	id := c.Param("id")

// 	var roleCreateOrEditDTO *backstagedto.RoleCreateOrEditDTO
// 	err = c.Bind(&roleCreateOrEditDTO)
// 	if err != nil {
// 		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
// 		log.Error(fmt.Sprintf("%+v", errData))
// 		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
// 	}
// 	roleService := backstage.GetRoleService()
// 	return roleService.EditRole(userInfo, id, roleCreateOrEditDTO)
// }
