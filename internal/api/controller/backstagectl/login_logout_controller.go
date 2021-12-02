package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils"
	"componentmod/internal/utils/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	BackstageLogin  = controller.Handler(Login)
	BackstageLogout = controller.Handler(Logout)
)

// @Summary Backstage login
// @Success 200 {json} json
// @param pwd path string true "pwd"
// @Router /backstage/login [post]
func Login(c *gin.Context) (controller.Data, error) {
	var loginDTO *backstagedto.LoginDTO
	err := c.Bind(&loginDTO)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	// backstageService.GetLoginLoutService()

	return nil, nil
}

// @Summary Backstage check password
// @Success 200 {json} json
// @param pwd path string true "pwd"
// @Router /backstage/login [post]
func Logout(c *gin.Context) (controller.Data, error) {
	// userService := backstage.GetUserService()
	return nil, nil
}
