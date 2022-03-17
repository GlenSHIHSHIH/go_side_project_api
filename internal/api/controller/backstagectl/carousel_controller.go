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
	CarouselShow    = controller.Handler(Carousels)
	CarouselIndex   = controller.Handler(CarouselById)
	CarouselDestory = controller.Handler(CarouselDelete)
	CarouselStore   = controller.Handler(CarouselCreate)
	CarouselUpdate  = controller.Handler(CarouselEdit)
)

//table list

// @tags Backstage-Carousel
// @Summary Carousel View
// @accept application/json
// @Success 200 {object} backstagedto.CarouselListDTO
// @Param page query int true "int default" default(1)
// @Param pageLimit query int true "int enums" Enums(15,20,30,40,50)
// @Param sort query string true "string enums" Enums(asc,desc)
// @Param sortColumn query string true "string enums" Enums(id,key)
// @Param search query string false "string default" default()
// @Param searchCategory query string false "string default" default()
// @Router /backstage/carousel [get]
func Carousels(c *gin.Context) (controller.Data, error) {
	search := c.QueryMap("search")
	var pageForMultSearchDTO = GetPageMultSearchDefaultDTO()
	pageForMultSearchDTO.Search = search

	err := c.Bind(pageForMultSearchDTO)
	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	CarouselService := backstage.GetCarouselService()
	return CarouselService.GetCarouselViewList(pageForMultSearchDTO)
}

// @tags Backstage-Carousel
// @Summary Carousel By Id
// @accept application/json
// @Success 200 {object} backstagedto.CarouselIdDTO
// @param id path int true "id"
// @Router /backstage/carousel/{id} [get]
func CarouselById(c *gin.Context) (controller.Data, error) {
	id := c.Param("id")

	CarouselService := backstage.GetCarouselService()
	return CarouselService.GetCarouselById(id)
}

// @tags Backstage-Carousel
// @Summary Carousel Delete
// @accept application/json
// @Success 200
// @param id path int true "id"
// @Router /backstage/carousel/delete/{id} [delete]
func CarouselDelete(c *gin.Context) (controller.Data, error) {
	ids := strings.Split(c.Param("id"), ",")

	roleService := backstage.GetRoleService()
	return roleService.DeleteRole(ids)
}

// @tags Backstage-Carousel
// @Summary Carousel Create
// @accept application/json
// @Success 200
// @Param json body backstagedto.CarouselCreateOrEditDTO true "json"
// @Router /backstage/carousel/create [post]
func CarouselCreate(c *gin.Context) (controller.Data, error) {

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

// @tags Backstage-Carousel
// @Summary Carousel Edit
// @accept application/json
// @Success 200
// @param id path int true "id"
// @Param json body backstagedto.CarouselCreateOrEditDTO true "json"
// @Router /backstage/carousel/edit/{id} [put]
func CarouselEdit(c *gin.Context) (controller.Data, error) {
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
