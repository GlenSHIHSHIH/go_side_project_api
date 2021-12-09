package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/validator.v2"
)

var (
	UserCreate = controller.Handler(CreateUser)
	UserEdit   = controller.Handler(EditUser)
)

// @tags Backstage
// @Summary Backstage UserLogin
// @accept application/json
// @produce application/json
// @Success 200 {object} dto.BaseResponseDTO
// @Param json body backstagedto.UserDTO true "json"
// @Router /backstage/user/create [post]
func CreateUser(c *gin.Context) (controller.Data, error) {
	var user *model.User
	err := c.Bind(&user)
	err = validator.Validate(user)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	userService := backstage.GetUserService()
	return userService.CreateUser(user)
}

func EditUser(c *gin.Context) (controller.Data, error) {
	JwtInfoDTO, _ := c.Get("userInfo")
	userInfo := JwtInfoDTO.(*backstagedto.JwtInfoDTO)
	return userInfo, nil
}
