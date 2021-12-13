package backstagectl

import (
	"componentmod/internal/api/controller"
	"componentmod/internal/api/errorcode"
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
	BackstageLogin        = controller.Handler(Login)
	BackstageLogout       = controller.Handler(Logout)
	BackstageRefreshToken = controller.Handler(RefreshToken)
	BackstageCheckToken   = controller.Handler(CheckToken)
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

	loginLogoutService := backstage.GetLoginLogoutService()

	return loginLogoutService.Login(loginDTO)
}

func Logout(c *gin.Context) (controller.Data, error) {
	// userService := backstage.GetUserService()
	return nil, nil
}

// @tags Backstage
// @Summary Backstage RefreshToken
// @accept application/json
// @produce application/json
// @Success 200 {object} backstagedto.LoginResponseDTO
// @Param json body backstagedto.JwtRefTokenDTO true "json"
// @Router /backstage/jwt/refreshtoken [post]
func RefreshToken(c *gin.Context) (controller.Data, error) {
	var jwtRefTokenDTO *backstagedto.JwtRefTokenDTO
	err := c.Bind(&jwtRefTokenDTO)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	loginLogoutService := backstage.GetLoginLogoutService()
	return loginLogoutService.RefreshToken(jwtRefTokenDTO)
}

// @tags Backstage
// @Summary Backstage CheckToken
// @accept application/json
// @produce application/json
// @Success 200 {object} backstagedto.JwtInfoDTO
// @Param string header string true "Authorization"
// @Router /backstage/jwt/check [post]
func CheckToken(c *gin.Context) (controller.Data, error) {
	bearerToken := c.GetHeader("Authorization")
	token := strings.Replace(bearerToken, "Bearer ", "", 1)

	//驗證使用者 jwt token
	jwtInfoDTO, err := utils.ValidateAndTokenCheck(token)

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.PARAMETER_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	return jwtInfoDTO, nil
}
