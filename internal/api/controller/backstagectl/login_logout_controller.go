package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/services/api/backstage"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/validator.v2"
)

var (
	BackstageLogin  = controller.Handler(Login)
	BackstageLogout = controller.Handler(Logout)
)

// @tags Backstage
// @Summary Backstage login
// @accept application/json
// @produce application/json
// @Success 200 {object} backstagedto.LoginResponseDTO
// @Param json body backstagedto.LoginDTO true "json"
// @Router /backstage/admin/login [post]
func Login(c *gin.Context) (controller.Data, error) {
	var loginDTO *backstagedto.LoginDTO
	err := c.Bind(&loginDTO)
	err = validator.Validate(loginDTO)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	loginLoutService := backstage.GetLoginLoutService()

	return loginLoutService.Login(loginDTO)
}

func Logout(c *gin.Context) (controller.Data, error) {
	// userService := backstage.GetUserService()
	return nil, nil
}
