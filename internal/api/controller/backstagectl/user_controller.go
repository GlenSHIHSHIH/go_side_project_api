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
	"gopkg.in/validator.v2"
)

var (
	UserShow      = controller.Handler(Users)
	UserIndex     = controller.Handler(UserById)
	UserStore     = controller.Handler(UserCreate)
	UserUpdate    = controller.Handler(UserEdit)
	UserPwdUpdate = controller.Handler(UserPwdEdit)
	UserPwdSet    = controller.Handler(UserPwdReset)
	UserDestory   = controller.Handler(UserDelete)
)

// @tags Backstage-User
// @Summary User Create
// @accept application/json
// @Security BearerAuth
// @Success 200
// @Param json body backstagedto.UserCreateOrEditDTO true "json"
// @Router /backstage/user/create [post]
func UserCreate(c *gin.Context) (controller.Data, error) {

	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	var userCreateOrEditDTO *backstagedto.UserCreateOrEditDTO
	err = c.Bind(&userCreateOrEditDTO)
	err = validator.Validate(userCreateOrEditDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	userService := backstage.GetUserService()
	return userService.CreateUser(userInfo, userCreateOrEditDTO)
}

// @tags Backstage-User
// @Summary User Delete
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Router /backstage/user/delete/{id} [delete]
func UserDelete(c *gin.Context) (controller.Data, error) {
	ids := strings.Split(c.Param("id"), ",")

	userService := backstage.GetUserService()
	return userService.DeleteUser(ids)
}

// @tags Backstage-User
// @Summary User Edit
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Param json body backstagedto.UserCreateOrEditDTO true "json"
// @Router /backstage/user/edit/{id} [put]
func UserEdit(c *gin.Context) (controller.Data, error) {
	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	id := c.Param("id")

	var userCreateOrEditDTO *backstagedto.UserCreateOrEditDTO
	err = c.Bind(&userCreateOrEditDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}
	userService := backstage.GetUserService()
	return userService.EditUser(userInfo, id, userCreateOrEditDTO)
}

// @tags Backstage-User
// @Summary User Passowrd Edit
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Param json body backstagedto.UserEditPwdDTO true "json"
// @Router /backstage/user/password/edit/{id} [put]
func UserPwdEdit(c *gin.Context) (controller.Data, error) {
	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	id := c.Param("id")

	var UserEditPwdDTO *backstagedto.UserEditPwdDTO
	err = c.Bind(&UserEditPwdDTO)
	err = validator.Validate(UserEditPwdDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}
	userService := backstage.GetUserService()
	return userService.EditUserPwd(userInfo, id, backstage.EDIT_TYPE, UserEditPwdDTO)
}

// @tags Backstage-User
// @Summary User Passowrd Reset
// @accept application/json
// @Security BearerAuth
// @Success 200
// @param id path int true "id"
// @Param json body backstagedto.UserEditPwdDTO true "json"
// @Router /backstage/user/password/reset/{id} [put]
func UserPwdReset(c *gin.Context) (controller.Data, error) {
	userInfo, err := validate.UserInfoValidate(c)
	if err != nil {
		return nil, err
	}

	id := c.Param("id")

	var UserEditPwdDTO *backstagedto.UserEditPwdDTO
	err = c.Bind(&UserEditPwdDTO)
	err = validator.Validate(UserEditPwdDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}
	userService := backstage.GetUserService()
	return userService.EditUserPwd(userInfo, id, backstage.RESET_TYPE, UserEditPwdDTO)
}

// @tags Backstage-User
// @Summary User By Id
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.UserIdDTO
// @param id path int true "id"
// @Router /backstage/user/{id} [get]
func UserById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")

	userService := backstage.GetUserService()
	return userService.GetUserById(id)
}

// @tags Backstage-User
// @Summary User View
// @accept application/json
// @Security BearerAuth
// @Success 200 {object} backstagedto.UserListDTO
// @Param page query int true "int default" default(1)
// @Param pageLimit query int true "int enums" Enums(15,20,30,40,50)
// @Param sort query string true "string enums" Enums(asc,desc)
// @Param sortColumn query string true "string enums" Enums(id,name,login_name)
// @Param search query string false "string default" default()
// @Param searchCategory query string false "string default" default()
// @Router /backstage/user [get]
func Users(c *gin.Context) (controller.Data, error) {
	search := c.QueryMap("search")
	var pageForMultSearchDTO = GetPageMultSearchDefaultDTO()
	pageForMultSearchDTO.Search = search

	err := c.Bind(pageForMultSearchDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	userService := backstage.GetUserService()
	return userService.GetUserViewList(pageForMultSearchDTO)
}
